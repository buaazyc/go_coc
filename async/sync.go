package async

import (
	"log"
	"time"

	"go_coc/constant"
	"go_coc/dao"
	"go_coc/goroutine"
	"go_coc/scene"
)

// Init 启动定时任务
func Init() error {
	if err := sync(); err != nil {
		return err
	}
	goroutine.GoWithRecover(func() {
		for range time.Tick(constant.SyncTime) {
			_ = sync()
		}
	})
	return nil
}

// sync 刷新任务都放在这里
func sync() error {
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
		if _, err := scene.CurrentWar(clan); err != nil {
			log.Printf("scene.CurrentWar err: %v", err)
			continue
		}
	}
	return nil
}
