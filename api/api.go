package api

import (
	"net/http"

	"go_coc/config"
	"go_coc/goroutine"
)

// Server 启动服务
func Server() error {
	// 启动服务处理https请求
	goroutine.GoWithRecover(func() {
		// 使用handler函数处理请求
		http.HandleFunc("/", addHeader(handler))
		// 监听特定端口，采用https加密传输
		http.ListenAndServeTLS(config.Conf.ServerPort, config.Conf.CertFile, config.Conf.KeyFile, nil)
	})
	select {}
}
