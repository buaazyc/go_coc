package scene

import (
	"go_coc/client"
	"go_coc/dao"
	"go_coc/parser"
)

func CurrentWar() error {
	res, err := client.SendAPI("/clans/%23R2JRG9PQ/currentwar")
	if err != nil {
		return err
	}
	currentWar, _ := parser.CurrentWar(res)
	if err := dao.Insert(currentWar.Clan.Tag, currentWar.StartTime, res); err != nil {
		return err
	}
	return nil
}
