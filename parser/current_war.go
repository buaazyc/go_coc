package parser

import (
	"encoding/json"
	"log"
)

func CurrentWar(info string) error {
	// log.Println(info)
	warInfo := make(map[string]interface{})
	err := json.Unmarshal([]byte(info), &warInfo)
	for k, v := range warInfo {
		log.Println(k)
		log.Println(v)
		log.Println()
	}
	if err != nil {
		return err
	}
	return nil
}
