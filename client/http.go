package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go_coc/config"
	"go_coc/constant"
)

// httpClient 超时时间设置为3s
var httpClient = &http.Client{
	Timeout: time.Second * 3,
}

// SendAPI 向coc开发者api发送http请求
func SendAPI(uri string) (string, error) {
	// 创建GET请求
	url := constant.CocBaseURL + uri
	log.Printf("send url: %v", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest err: %v", err)
	}
	req.Header.Add("authorization", "Bearer "+config.Conf.Token)
	req.Header.Add("Accept", "application/json")
	// 发起请求
	rsp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("httpClient.Do err: %v", err)
	}
	defer rsp.Body.Close()
	// 获取内容
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	return string(body), nil
}
