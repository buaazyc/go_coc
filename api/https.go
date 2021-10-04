// Package api 存放前后端通信接口相关
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go_coc/constant"
	"go_coc/scene"
)

// handler 服务监听函数
func handler(w http.ResponseWriter, req *http.Request) {
	// 解析客户端发送的req请求内容
	req.ParseForm()
	use := req.FormValue("use")
	clan := req.FormValue("clan")
	if len(clan) < constant.MinClanLen {
		return
	}
	clan = strings.ToTitle(clan)
	log.Printf("use: %v clan: %v", use, clan)
	// 根据use不同，触发不同的场景
	switch use {
	case "currentwar":
		cur, err := scene.CurrentWar(clan[1:])
		if err != nil {
			log.Printf("cache.CurrentWar err: %v", err)
			errRsp(w, 404)
			return
		}
		res, err := json.Marshal(cur)
		if err != nil {
			log.Printf("json.Marshal err: %v", err)
			errRsp(w, 404)
			return
		}
		fmt.Fprintf(w, "%+v", string(res))
	case "leaguegroup":
		group, err := scene.LeagueGroup(clan[1:])
		if err != nil {
			log.Printf("scene.LeagueGroup err: %v", err)
			errRsp(w, 404)
			return
		}
		groupRsp, err := scene.LeagueGroupRsp(group)
		if err != nil {
			log.Printf("scene.LeagueGroupRsp err: %v", err)
			errRsp(w, 404)
			return
		}
		res, err := json.Marshal(groupRsp)
		if err != nil {
			log.Printf("json.Marshal err: %v", err)
			errRsp(w, 404)
			return
		}
		fmt.Fprintf(w, "%v", string(res))
	case "leaguewar":
		leaguewar, err := scene.LeagueWarRsp(clan[1:])
		if err != nil {
			log.Printf("scene.LeagueWarRsp err: %v", err)
			errRsp(w, 404)
			return
		}
		res, err := json.Marshal(leaguewar)
		if err != nil {
			log.Printf("json.Marshal err: %v", err)
			errRsp(w, 404)
			return
		}
		fmt.Fprintf(w, "%+v", string(res))
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

type ErrRsp struct {
	Constructor uint32 `json:"constructor"`
}

func errRsp(w http.ResponseWriter, errCode uint32) {
	res, _ := json.Marshal(ErrRsp{Constructor: errCode})
	fmt.Fprintf(w, "%+v", string(res))
}
