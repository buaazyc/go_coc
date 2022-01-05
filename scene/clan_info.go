package scene

import (
	"fmt"
	"net/http"

	"go_coc/client"
	"go_coc/constant"
	"go_coc/parser"
)

type MembersScene struct{}

func init() {
	register(constant.MembersScene, &MembersScene{})
}

func (s *MembersScene) Do(clan string, w http.ResponseWriter) error {
	clanInfo, err := clanInfo(clan)
	if err != nil {
		return err
	}
	if clanInfo.MemberList == nil {
		return fmt.Errorf("clanInfo.MemberList == nil")
	}
	return response(w, clanInfo.MemberList)
}

func clanInfo(clan string) (*parser.Clan, error) {
	// 向官方发送请求，获取最新数据
	res, err := client.SendAPI("/clans/%23" + clan[1:])
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
