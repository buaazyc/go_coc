package scene

import (
	"go_coc/client"
	"go_coc/parser"
	"log"
)

func PlayerInfo(player string) (*parser.Player, error) {
	res, err := client.SendAPI("/players/%23" + player[1:])
	if err != nil {
		return nil, err
	}
	playerInfo, err := parser.PlayerInfo(res)
	if err != nil {
		return nil, err
	}
	log.Printf("%+v", playerInfo)
	return playerInfo, nil
}
