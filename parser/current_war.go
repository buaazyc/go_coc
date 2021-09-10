package parser

import (
	"encoding/json"
)

// CurrentWar 解析部落战信息
func CurrentWar(info string) (*ClanWar, error) {
	warInfo := ClanWar{}
	err := json.Unmarshal([]byte(info), &warInfo)
	if err != nil {
		return nil, err
	}
	return &warInfo, nil
}
