package scene

import (
	"go_coc/client"
	"go_coc/dao"
	"go_coc/parser"
	"log"
)

// CurrentWar 获取当前部落战
func CurrentWar(clan string) (*parser.ClanWar, error) {
	res, err := client.SendAPI("/clans/%23" + clan + "/currentwar")
	if err != nil {
		return nil, err
	}
	currentWar, _ := parser.CurrentWar(res)
	log.Printf("currentWar: %+v", currentWar)
	if err := dao.InsertCurrentWar(currentWar.Clan.Tag, currentWar.StartTime, res); err != nil {
		return nil, err
	}
	return currentWar, nil
}
