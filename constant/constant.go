package constant

import "time"

// CocBaseURL coc官方api
const CocBaseURL = "https://api.clashofclans.com/v1"

// MinClanLen 允许接受的最短clan长度
const MinClanLen = 2

// HttpClientTimeOut 最为客户端发送http请求时，超时时间，单位s
const HttpClientTimeOut = 10

// SyncTime 后台定时任务都刷新周期
const (
	SyncLowFrequencyTime  = time.Minute * 10
	SyncHighFrequencyTime = time.Minute
)

// MyClan 本部落标签
const MyClan = "#R2JRG9PQ"

// 场景类型
const (
	CurrentWarScene  = "currentwar"
	LeaguegroupScene = "leaguegroup"
	LeaguewarScene   = "leaguewar"
	SeasonScene      = "season"
	MembersScene     = "members"
)
