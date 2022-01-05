package scene

import (
	"fmt"
	"net/http"

	"go_coc/client"
	"go_coc/constant"
	"go_coc/dao"
	"go_coc/parser"
)

type CurrentWarScene struct{}

func init() {
	register(constant.CurrentWarScene, &CurrentWarScene{})
}

func (s *CurrentWarScene) Do(clan string, w http.ResponseWriter) error {
	// 从数据库中获取最新的部落战
	war, err := dao.QueryNewestWar(clan)
	if err != nil {
		return err
	}
	if war == nil {
		return fmt.Errorf("war is nil")
	}
	// 解析数据
	currentWar, err := parser.CurrentWar(war.Info)
	if err != nil {
		return err
	}
	// 返回结果
	return response(w, currentWar)
}

// UpdateCurrentWar 更新当前部落战
func UpdateCurrentWar(clan string) (*parser.ClanWar, error) {
	// 向官方发送请求，获取最新数据
	res, err := client.SendAPI("/clans/%23" + clan[1:] + "/currentwar")
	if err != nil {
		return nil, err
	}
	// 解析数据
	currentWar, err := parser.CurrentWar(res)
	if err != nil {
		return nil, err
	}
	// 插入数据到数据库
	if currentWar.Clan == nil {
		return nil, fmt.Errorf("currentWar.Clan[:%v] == nil", clan)
	}
	if err := dao.InsertCurrentWar(currentWar.Clan.Tag, currentWar.StartTime, res); err != nil {
		return nil, err
	}
	return currentWar, nil
}
