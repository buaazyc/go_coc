package scene

import (
	"go_coc/client"
	"go_coc/parser"
)

func ClanInfo(clan string) (*parser.Clan, error) {
	// 向官方发送请求，获取最新数据
	res, err := client.SendAPI("/clans/%23" + clan)
	if err != nil {
		return nil, err
	}
	// 解析数据
	info, err := parser.ClanInfo(res)
	if err != nil {
		return nil, err
	}
	// log.Printf("%+v", info)
	return info, nil
}
