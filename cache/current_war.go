package cache

import (
	"go_coc/dao"
	"go_coc/parser"
)

// CurrentWar 从缓存中读取CurrentWar
func CurrentWar(clan string) (*parser.ClanWar, error) {
	res, err := queryLastWar(clan)
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

// queryLastWar 从数据库中获取最新的部落战
func queryLastWar(clan string) (string, error) {
	wars, err := dao.QueryAllWarsFor(clan)
	if err != nil {
		return "", nil
	}
	last := &dao.CurrentWar{}
	for _, war := range wars {
		if last.Time < war.Time {
			last = war
		}
	}
	return last.Info, nil
}
