package main

import (
	"log"

	"go_coc/api"
	"go_coc/config"
	"go_coc/dao"
	"go_coc/scene"
)

func develop() {
	scene.SumWar("#R2JRG9PQ", "202109")
}

func main() {
	// 初始化log输出格式
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// 初始化获取配置文件
	if err := config.GetConf(); err != nil {
		log.Printf("GetConf err: %v", err)
	}
	// 初始化数据库
	if err := dao.ConnectDB(); err != nil {
		log.Printf("ConnectDB err: %v", err)
	}
	// 异步定时任务
	// if err := async.Init(); err != nil {
	// 	log.Printf("async.Init err: %v", err)
	// }
	// 运行开发时的临时代码
	develop()
	// 启动服务
	if err := api.Server(); err != nil {
		log.Fatal(err)
	}
}
