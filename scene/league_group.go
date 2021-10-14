package scene

import (
	"fmt"
	"sort"
	"strconv"
	"sync"

	"go_coc/client"
	"go_coc/goroutine"
	"go_coc/parser"
	"go_coc/time"
)

// LeagueGroup 获取部落联赛分组
func LeagueGroup(clan string) (*parser.ClanWarLeagueGroup, error) {
	// 向官方发送请求，获取最新数据
	res, err := client.SendAPI("/clans/%23" + clan[1:] + "/currentwar/leaguegroup")
	if err != nil {
		return nil, err
	}
	// 解析数据
	leagueGroup, err := parser.LeagueGroup(res)
	if err != nil {
		return nil, err
	}
	if leagueGroup.Season == "" {
		return nil, fmt.Errorf("leaguegroup not found")
	}
	return leagueGroup, nil
}

// LeagueGroupRsp 向前端展示小组分组的情况
func LeagueGroupRsp(leagueGroup *parser.ClanWarLeagueGroup) (*parser.LeagueGroupRsp, error) {
	if leagueGroup == nil {
		return nil, fmt.Errorf("leagueGroup == nil")
	}
	// 每个部落获取其基本信息
	var (
		resClans  [][]string
		warLeague string
		lock      sync.Mutex
		funcs     []func() error
	)
	for _, c := range leagueGroup.Clans {
		clan := c
		funcs = append(funcs, func() error {
			clanInfo, err := ClanInfo(clan.Tag)
			if clanInfo.WarLeague == nil {
				return fmt.Errorf("clanInfo.WarLeague == nil")
			}
			warLeague = clanInfo.WarLeague.Name
			if err != nil {
				return err
			}
			lock.Lock()
			defer lock.Unlock()
			resClans = append(resClans, []string{
				clan.Name,
				clan.Tag,
				fmt.Sprint(clan.ClanLevel),
				fmt.Sprint(clanInfo.WarWinStreak),
				fmt.Sprint(clanInfo.WarWins),
				fmt.Sprint(clanInfo.ClanPoints),
				fmt.Sprint(countTownHallLevel(clan)[14]),
			})
			return nil
		})
	}
	if err := goroutine.GoAndWait(funcs...); err != nil {
		return nil, err
	}
	return &parser.LeagueGroupRsp{
		League: warLeague,
		Season: time.SeasonStr(leagueGroup.Season),
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
