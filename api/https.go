// Package api 存放接口相关代码
package api

import (
	"encoding/json"
	"fmt"
	"go_coc/constant"
	"go_coc/scene"
	"log"
	"net/http"
	"strings"
)

// handler 服务监听函数
func handler(w http.ResponseWriter, req *http.Request) {
	// 校验和解析
	if req.Method != "GET" {
		log.Printf("Method[%+v] is not GET", req.Method)
		return
	}
	use := req.FormValue("use")
	clan := req.FormValue("clan")
	if len(clan) < constant.MinClanLen {
		return
	}
	// 标签转大写
	clan = strings.ToTitle(clan)
	log.Printf("use: %v clan: %v", use, clan)
	// 获取场景
	s, ok := scene.GetScene(use)
	if !ok {
		errRsp(w, 404)
	}
	// Do
	if err := s.Do(clan, w); err != nil {
		log.Printf("do err: %v", err)
		errRsp(w, 404)
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

type ErrRsp struct {
	ErrCode uint32 `json:"errCode"`
}

func errRsp(w http.ResponseWriter, errCode uint32) {
	res, _ := json.Marshal(ErrRsp{ErrCode: errCode})
	fmt.Fprintf(w, "%+v", string(res))
}

func reply(w http.ResponseWriter, v interface{}) {
	res, err := json.Marshal(v)
	if err != nil {
		log.Printf("json.Marshal err: %v", err)
		errRsp(w, 404)
		return
	}
	fmt.Fprintf(w, "%+v", string(res))
}
