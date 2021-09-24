package main

import (
	"log"

	"go_coc/api"
	"go_coc/config"
	"go_coc/dao"
)

func main() {
	// 获取配置文件
	if err := config.GetConf(); err != nil {
		log.Printf("GetConf err: %v", err)
	}
	// 连接数据库
	if err := dao.ConnectDB(); err != nil {
		log.Printf("ConnectDB err: %v", err)
	}
	// 启动服务
	if err := api.Server(); err != nil {
		log.Fatal(err)
	}
}
