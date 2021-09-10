package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration 配置
type Configuration struct {
	Token string `json:"token"`
}

// 全局配置变量
var Conf Configuration

// GetConf 获取配置文件中的配置信息
func GetConf() error {
	// 打开配置文件config.json
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("打开配置文件config.json错误：%v", err)
	}
	defer file.Close()
	// 解析json文件
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Conf); err != nil {
		return fmt.Errorf("decoder.Decode err: %v", err)
	}
	return nil
}
