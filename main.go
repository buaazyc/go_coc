package main

import (
	"log"

	"go_coc/config"
	"go_coc/dao"
	"go_coc/scene"
)

func main() {
	if err := config.GetConf(); err != nil {
		log.Fatal(err)
	}
	if err := dao.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	if err := scene.CurrentWar(); err != nil {
		log.Fatal(err)
	}
}
