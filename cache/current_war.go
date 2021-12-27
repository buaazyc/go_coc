package cache

import (
	"go_coc/dao"
	"go_coc/parser"
)

// CurrentWar 从缓存中读取CurrentWar
func CurrentWar(clan string) (*parser.ClanWar, error) {
	// 从数据库中获取最新的部落战
	war, err := dao.QueryNewestWar(clan)
	if err != nil {
		return nil, err
	}
	if war == nil {
		return nil, nil
	}

	// 解析数据
	currentWar, err := parser.CurrentWar(war.Info)
	if err != nil {
		return nil, err
	}
	return currentWar, nil
}
