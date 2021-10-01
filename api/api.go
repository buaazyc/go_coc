// Package api 存放前后端通信接口相关
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go_coc/cache"
	"go_coc/config"
	"go_coc/parser"
	"go_coc/scene"
)

// Server 启动服务
func Server() error {
	// 使用handler函数处理请求
	http.HandleFunc("/", addHeader(handler))
	// 监听特定端口，采用https加密传输
	if err := http.ListenAndServeTLS(config.Conf.ServerPort, config.Conf.CertFile, config.Conf.KeyFile, nil); err != nil {
		return err
	}
	return nil
}

// handler 服务监听函数
func handler(w http.ResponseWriter, req *http.Request) {
	// 解析客户端发送的req请求内容
	req.ParseForm()
	use := req.FormValue("use")
	clan := req.FormValue("clan")
	if len(clan) < 2 {
		return
	}
	// 根据use不同，触发不同的场景
	switch use {
	case "currentwar":
		var cur *parser.ClanWar
		// 读缓存，若不存在，则读官方api
		cur, _ = cache.CurrentWar(clan[1:])
		if cur == nil {
			cur, _ = scene.CurrentWar(clan[1:])
		}
		// 转化为小写json传输给前端
		res, err := json.Marshal(cur)
		if err != nil {
			log.Printf("err: %v", err)
			return
		}
		fmt.Fprintf(w, "%+v", strings.ToLower(string(res)))
	default:
	}
}

func addHeader(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")               // 允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")   // header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")       // 设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "GET")            // 允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8") // 返回数据格式是json
		f(w, r)
	}
}
