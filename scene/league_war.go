package scene

import (
	"fmt"
	"log"

	"go_coc/client"
	"go_coc/parser"
)

// LeagueWarRsp 汇总联赛战绩
func LeagueWarRsp(clan string) (*parser.LeagueWarRsp, error) {
	// 获取联赛小组全部信息
	clanWarLeagueGroup, err := LeagueGroup(clan)
	if err != nil {
		return nil, err
	}
	// 获取联赛中有效战绩
	wars, err := getValidWar(clan, clanWarLeagueGroup)
	if err != nil {
		return nil, err
	}
	if len(wars) == 0 {
		return nil, fmt.Errorf("clan[%v] len(wars) == 0", clan)
	}
	// 查询部落基本信息
	clanInfo, err := ClanInfo(clan)
	if err != nil {
		return nil, err
	}
	// 提取各成员战绩
	members := getMembers(wars)
	return &parser.LeagueWarRsp{
		Name:    wars[0].Name,
		Season:  seasonStr(clanWarLeagueGroup.Season),
		League:  clanInfo.WarLeague.Name,
		Members: members,
	}, nil
}

// LeagueWar 根据warTag获取战绩
func LeagueWar(war string) (*parser.ClanWar, error) {
	// 向官方发送请求，获取最新数据
	res, err := client.SendAPI("/clanwarleagues/wars/%23" + war)
	if err != nil {
		return nil, err
	}
	// 解析数据
	currentWar, err := parser.CurrentWar(res)
	if err != nil {
		return nil, err
	}
	return currentWar, nil
}

// getValidWar 获取clan参与的联赛战绩
func getValidWar(clan string, clanWarLeagueGroup *parser.ClanWarLeagueGroup) ([]*parser.WarClan, error) {
	var wars []*parser.WarClan
	// assert len(clanWarLeagueGroup.Rounds) == 7
	// 一次联赛打7天
	for _, round := range clanWarLeagueGroup.Rounds {
		// assert len(round.WarTags) == 4
		// 每天8个队打4场
		for _, warTag := range round.WarTags {
			// 未到准备日打warTag为 #0
			if warTag == "#0" {
				break
			}
			war, err := LeagueWar(warTag[1:])
			if err != nil {
				log.Printf("%v", err)
				continue
			}
			// 准备日的进攻不需要统计
			if war.State == "preparation" {
				continue
			}
			// 保存本部落相关战绩
			if war.Clan.Tag[1:] == clan {
				wars = append(wars, &war.Clan)
			}
			if war.Opponent.Tag[1:] == clan {
				wars = append(wars, &war.Opponent)
			}
		}
	}
	return wars, nil
}

// setStar 根据stars，在对应字段增加计数
func setStar(stars uint32, info *parser.AttackInfo) {
	if stars == 0 {
		info.Zero++
	}
	if stars == 1 {
		info.One++
	}
	if stars == 2 {
		info.Two++
	}
	if stars == 3 {
		info.Three++
	}
}

// getMembers 根据每一场战绩wars，提取各个成员战绩
func getMembers(wars []*parser.WarClan) map[string]*parser.LeagueWarMember {
	members := make(map[string]*parser.LeagueWarMember)
	// 遍历每一场
	for _, war := range wars {
		// 遍历每一场的每一个成员
		for _, member := range war.Members {
			// 如果是初次读取数据，需要初始化成员变量
			if members[member.Tag] == nil {
				members[member.Tag] = &parser.LeagueWarMember{}
				members[member.Tag].AttackInfo = &parser.AttackInfo{}
				members[member.Tag].Defend = &parser.Defend{}
				members[member.Tag].Name = member.Name
			}
			// 遍历每个成员的进攻
			for _, attack := range member.Attacks {
				setStar(attack.Stars, members[member.Tag].AttackInfo)
				members[member.Tag].AttackInfo.AttackNum++
				members[member.Tag].AttackInfo.SumStars += attack.Stars
				members[member.Tag].AttackInfo.SumDestruction += attack.DestructionPercentage
				members[member.Tag].AttackInfo.SumDuration += attack.Duration
			}
			// 总场次
			members[member.Tag].JoinNum++
			// 防守情况
			members[member.Tag].Defend.SumStars += member.BestOpponentAttack.Stars
			if member.BestOpponentAttack.Stars == 3 {
				members[member.Tag].Defend.Three++
			}
		}
	}
	// 计算每个成员的得分
	for _, member := range members {
		member.Score = calScore(member)
	}
	return members
}

// calScore 计算每个成员的得分情况
func calScore(member *parser.LeagueWarMember) uint32 {
	var score uint32
	score += 30 * (member.AttackInfo.AttackNum / member.JoinNum)
	score += 15 * (member.AttackInfo.Three / member.JoinNum)
	score += 15 * ((member.JoinNum - member.AttackInfo.One) / member.JoinNum)
	score += 20 * ((member.JoinNum - member.AttackInfo.Zero) / member.JoinNum)
	score += 20 * ((member.JoinNum - member.Defend.Three) / member.JoinNum)
	return score
}
