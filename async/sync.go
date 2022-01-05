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
func Init() {
	// 低优先级
	goroutine.GoWithRecover(func() {
		for range time.Tick(constant.SyncLowFrequencyTime) {
			cacheActiveCurrentWar()
		}
	})
	// 高优先级
	goroutine.GoWithRecover(func() {
		for range time.Tick(constant.SyncHighFrequencyTime) {
			cacheMyCurrentWar()
		}
	})
}

// cacheActiveCurrentWar 缓存活跃部落的部落战战绩
func cacheActiveCurrentWar() {
	// 获取活跃部落标签
	clans, err := dao.QueryActiveClanTags()
	if err != nil {
		log.Printf("dao.QueryActiveClanTags err: %v", err)
	}
	log.Printf("clans number from dao.QueryAllClanTags is: %v", len(clans))
	// 遍历所有部落标签，查询最新战绩
	for _, clan := range clans {
		if _, err := scene.UpdateCurrentWar(clan); err != nil {
			log.Printf("scene.CurrentWar err: %v", err)
			continue
		}
	}
}

// cacheMyCurrentWar 缓存本部落战绩
func cacheMyCurrentWar() {
	_, _ = scene.UpdateCurrentWar(constant.MyClan)
}
