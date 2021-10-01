package cache

import (
	"go_coc/dao"
	"go_coc/parser"
)

// CurrentWar 从缓存中读取CurrentWar
func CurrentWar(clan string) (*parser.ClanWar, error) {
	res, err := dao.QueryLastCurrentWar(clan)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, err
	}
	// 解析数据
	currentWar, err := parser.CurrentWar(res)
	if err != nil {
		return nil, err
	}
	return currentWar, nil
}
