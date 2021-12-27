package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration 配置
type Configuration struct {
	// coc开发者配置
	Token string `json:"token"`
	// mysql配置
	MysqlUser     string `json:"mysqlUser"`
	MysqlPassword string `json:"mysqlPassword"`
	MysqlDBName   string `json:"mysqlDBName"`
	MysqlHost     string `json:"mysqlHost"`
	MysqlPort     int    `json:"mysqlPort"`
	// http配置
	ServerPort string `json:"serverPort"`
	CertFile   string `json:"certFile"`
	KeyFile    string `json:"keyFile"`
}

// 全局配置变量
var Conf Configuration

// GetConf 获取配置文件中的配置信息
func GetConf() error {
	// 打开配置文件config.json
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("open config.json err: %v", err)
	}
	defer file.Close()

	// 解析json文件
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Conf); err != nil {
		return fmt.Errorf("decoder.Decode err: %v", err)
	}
	return nil
}
