package scene

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"go_coc/client"
	"go_coc/parser"
)

// LeagueGroup 获取部落联赛分组
func LeagueGroup(clan string) (*parser.ClanWarLeagueGroup, error) {
	// 向官方发送请求，获取最新数据
	res, err := client.SendAPI("/clans/%23" + clan + "/currentwar/leaguegroup")
	if err != nil {
		return nil, err
	}
	// 解析数据
	leagueGroup, err := parser.LeagueGroup(res)
	if err != nil {
		return nil, err
	}
	return leagueGroup, nil
}

// LeagueGroupRsp 在前端显示联赛分组信息
func LeagueGroupRsp(leagueGroup *parser.ClanWarLeagueGroup) (*parser.LeagueGroupRsp, error) {
	if leagueGroup == nil {
		return nil, fmt.Errorf("leagueGroup == nil")
	}
	// 每个部落获取其基本信息
	var resClans [][]string
	var warLeague string
	for _, clan := range leagueGroup.Clans {
		townHall := countTownHallLevel(&clan)
		clanInfo, err := ClanInfo(clan.Tag[1:])
		warLeague = clanInfo.WarLeague.Name
		if err != nil {
			log.Printf("get ClanInfo err : %v", err)
			continue
		}
		resClans = append(resClans, []string{
			clan.Name,
			clan.Tag,
			fmt.Sprint(clan.ClanLevel),
			fmt.Sprint(clanInfo.WarWinStreak),
			fmt.Sprint(clanInfo.WarWins),
			fmt.Sprint(clanInfo.ClanPoints),
			fmt.Sprint(townHall[14]),
		})
	}
	return &parser.LeagueGroupRsp{
		League: warLeague,
		Season: seasonStr(leagueGroup.Season),
		Clans:  sortClans(resClans),
	}, nil
}

// countTownHallLevel 统计参战的大本等级统计
func countTownHallLevel(clan *parser.ClanWarLeagueClan) map[uint32]uint32 {
	res := make(map[uint32]uint32)
	for _, member := range clan.Members {
		res[member.TownHallLevel]++
	}
	return res
}

// seasonStr 将赛季的年月以汉字形式输出
func seasonStr(s string) string {
	return fmt.Sprintf("%v年%v月", s[:4], s[5:])
}

// sortClans 将部落排序
func sortClans(clans [][]string) [][]string {
	// 按照部落等级排序
	sort.Slice(clans, func(i, j int) bool {
		a, _ := strconv.Atoi(clans[i][2])
		b, _ := strconv.Atoi(clans[j][2])
		return a > b
	})
	return clans
}
