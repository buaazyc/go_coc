// Package api 存放前后端通信接口相关
package api

import (
	"fmt"
	"net/http"

	"go_coc/config"
	"go_coc/scene"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	use := req.FormValue("use")
	clan := req.FormValue("clan")
	switch use {
	case "currentwar":
		cur, _ := scene.CurrentWar(clan[1:])
		fmt.Fprintf(w, "%+v", cur)
	default:
	}
}

func Server() error {
	http.HandleFunc("/", indexHandler)
	return http.ListenAndServeTLS(config.Conf.ServerPort, config.Conf.CertFile, config.Conf.KeyFile, nil)
}
