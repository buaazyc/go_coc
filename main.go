package main

import (
	"log"

	"go_coc/client"
	"go_coc/config"
	"go_coc/parser"
)

func main() {
	if err := config.GetConf(); err != nil {
		log.Fatal(err)
	}
	res, err := client.SendAPI("/clans")
	// res, err := client.SendAPI("/clans/%23R2JRG9PQ/currentwar")
	if err != nil {
		log.Fatal(err)
	}
	parser.CurrentWar(res)

}
