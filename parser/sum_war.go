package parser

type LeagueWarRsp struct {
	Name    string                `json:"name"`
	Season  string                `json:"season"`
	League  string                `json:"league"`
	Members map[string]*WarMember `json:"members"`
}

type SeasonRsp struct {
	Name    string                `json:"name"`
	Season  string                `json:"season"`
	Members map[string]*WarMember `json:"members"`
}

type WarMember struct {
	Name       string      `json:"name"`
	AttackInfo *AttackInfo `json:"attackInfo"`
	Defend     *Defend     `json:"defend"`
	JoinNum    uint32      `json:"joinNum"`
	Score      uint32      `json:"score"`
}

type AttackInfo struct {
	Three          uint32 `json:"three"`
	Two            uint32 `json:"two"`
	One            uint32 `json:"one"`
	Zero           uint32 `json:"zero"`
	AttackNum      uint32 `json:"attackNum"`
	SumStars       uint32 `json:"sumStars"`
	SumDestruction uint32 `json:"sumDestruction"`
	SumDuration    uint32 `json:"sumDuration"`
}

type Defend struct {
	Three    uint32 `json:"three"`
	SumStars uint32 `json:"sumStars"`
}
