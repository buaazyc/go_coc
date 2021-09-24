// Package api 存放前后端通信接口相关
package api

import (
	"fmt"
	"net/http"

	"go_coc/config"
	"go_coc/scene"
)

// handler 服务监听函数
func handler(w http.ResponseWriter, req *http.Request) {
	// 解析客户端发送的req请求内容
	req.ParseForm()
	use := req.FormValue("use")
	clan := req.FormValue("clan")
	// 根据use不同，触发不同的场景
	switch use {
	case "currentwar":
		cur, _ := scene.CurrentWar(clan[1:])
		fmt.Fprintf(w, "%+v", cur)
	default:
	}
}

// Server 启动服务
func Server() error {
	http.HandleFunc("/", handler)
	return http.ListenAndServeTLS(config.Conf.ServerPort, config.Conf.CertFile, config.Conf.KeyFile, nil)
}
