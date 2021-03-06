package main

import (
	"log"

	"go_coc/api"
	"go_coc/async"
	"go_coc/config"
	"go_coc/dao"
)

func develop() {
}

func main() {
	// 运行开发时的临时代码
	develop()
	// 初始化log输出格式
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// 初始化获取配置文件
	if err := config.GetConf(); err != nil {
		log.Fatalf("GetConf err: %v", err)
	}
	// 初始化数据库
	if err := dao.ConnectDB(); err != nil {
		log.Fatalf("ConnectDB err: %v", err)
	}
	// 异步定时任务
	async.Init()

	// 启动服务
	if err := api.Server(); err != nil {
		log.Fatalf("server err: %v", err)
	}
}
