package scene

import (
	"fmt"
	"log"

	"go_coc/dao"
	"go_coc/parser"
	"go_coc/time"
)

// Season 部落战赛季汇总
func Season(clan string, season string) (*parser.SeasonRsp, error) {
	monthWars, err := queryMonthWars(clan, season)
	if err != nil {
		return nil, err
	}
	if len(monthWars) == 0 {
		return nil, fmt.Errorf("len(monthWars) == 0")
	}
	members := getMembers(monthWars)
	calSeasonScore(members)
	return &parser.SeasonRsp{
		Name:    monthWars[0].Name,
		Season:  time.SeasonToStr(season),
		Members: members,
	}, nil
}

// CurSeason 当前赛季的赛季汇总
func CurSeason(clan string) (*parser.SeasonRsp, error) {
	return Season(clan, time.GetCurSeason())
}

// queryMonthWars 从数据库中获取部落整个月的战绩
func queryMonthWars(clan string, season string) ([]*parser.WarClan, error) {
	wars, err := dao.QueryAllWars(clan)
	if err != nil {
		return nil, err
	}
	var res []*parser.WarClan
	for _, war := range wars {
		monthTime, err := time.TimeToMonth(war.Time)
		if err != nil {
			log.Print(err)
			continue
		}
		if monthTime == season {
			// 解析数据
			parsedWar, err := parser.CurrentWar(war.Info)
			if err != nil {
				return nil, err
			}
			res = append(res, parsedWar.Clan)
		}
	}
	return res, nil
}

// calSeasonScore 计算每个成员部落战汇总的得分情况
func calSeasonScore(members map[string]*parser.WarMember) {
	for _, member := range members {
		var score uint32
		score += 30 * (member.AttackInfo.AttackNum / (2 * member.JoinNum))
		score += 15 * (member.AttackInfo.Three / (2 * member.JoinNum))
		score += 15 * ((member.JoinNum - member.AttackInfo.One) / member.JoinNum)
		score += 20 * ((member.JoinNum - member.AttackInfo.Zero) / member.JoinNum)
		score += 20 * ((member.JoinNum - member.Defend.Three) / member.JoinNum)
		member.Score = score
	}
}
