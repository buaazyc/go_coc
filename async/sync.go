package async

import (
	"go_coc/dao"
	"go_coc/scene"
	"log"
	"time"
)

// Init 启动定时任务
func Init() error {
	if err := sync5Min(); err != nil {
		return err
	}
	GoWithRecover(func() {
		for range time.Tick(time.Minute * 5) {
			_ = sync5Min()
		}
	})
	return nil
}

// sync5Min 每5分钟刷新一次
func sync5Min() error {
	if err := currentWar(); err != nil {
		return err
	}
	return nil
}

// currentWar 缓存currentWar
func currentWar() error {
	clans, err := dao.QueryAllClanTags()
	if err != nil {
		return err
	}
	log.Printf("clans number from dao.QueryAllClanTags is: %v", len(clans))
	for _, clan := range clans {
		if len(clan) < 2 {
			continue
		}
		if _, err := scene.CurrentWar(clan[1:]); err != nil {
			log.Printf("scene.CurrentWar err: %v", err)
			continue
		}
	}
	return nil
}