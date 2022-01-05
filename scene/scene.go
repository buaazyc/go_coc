package scene

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Scene 场景接口
type Scene interface {
	// Do 查询部落特定场景的数据，并写到ResponseWriter
	Do(clan string, w http.ResponseWriter) error
}

// GetScene 获取特定接口
func GetScene(use string) (Scene, bool) {
	v, ok := sceneMap[use]
	return v, ok
}

var sceneMap = make(map[string]Scene)

// register 注册
func register(t string, s Scene) {
	sceneMap[t] = s
}

// response 返回数据
func response(w http.ResponseWriter, v interface{}) error {
	res, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json.Marshal err: %v", err)
	}
	fmt.Fprintf(w, "%+v", string(res))
	return nil
}
