package scene

import (
	"go_coc/cache"
	"go_coc/client"
	"go_coc/dao"
	"go_coc/parser"
	"log"
)

// CurrentWar 获取当前部落战
func CurrentWar(clan string) (*parser.ClanWar, error) {
	// 读缓存，若不存在，则读官方api
	cur, err := cache.CurrentWar(clan[1:])
	if err != nil {
		log.Printf("cache.CurrentWar err: %v", err)
		return nil, err
	}
	if cur != nil {
		return cur, nil
	}
	// 向官方发送请求，获取最新数据
	res, err := client.SendAPI("/clans/%23" + clan + "/currentwar")
	if err != nil {
		return nil, err
	}
	// 解析数据
	currentWar, err := parser.CurrentWar(res)
	if err != nil {
		return nil, err
	}
	// 插入数据到数据库
	if err := dao.InsertCurrentWar(currentWar.Clan.Tag, currentWar.StartTime, res); err != nil {
		return nil, err
	}
	return currentWar, nil
}
