package api

import (
	"go_coc/async"
	"go_coc/config"
	"net/http"
)

// Server 启动服务
func Server() error {
	// 启动服务处理https请求
	async.GoWithRecover(func() {
		// 使用handler函数处理请求
		http.HandleFunc("/", addHeader(handler))
		// 监听特定端口，采用https加密传输
		http.ListenAndServeTLS(config.Conf.ServerPort, config.Conf.CertFile, config.Conf.KeyFile, nil)
	})
	select {}
}
