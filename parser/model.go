package parser

// ClanWar 当前部落战信息
type ClanWar struct {
	Clan                 WarClan
	Opponent             WarClan
	TeamSize             uint32
	StartTime            string
	EndTime              string
	State                string
	PreparationStartTime string
}

// WarClan 参战部落信息
type WarClan struct {
	DestructionPercentage float64
	Tag                   string
	Name                  string
	BadgeUrls             BadgeUrl
	ClanLevel             uint32
	Attacks               uint32
	Stars                 uint32
	ExpEarned             uint32
	Members               []ClanWarMember
}

// BadgeUrl 部落图标
type BadgeUrl struct {
	Large  string
	Medium string
	Samll  string
}

// ClanWarMember 参战成员信息
type ClanWarMember struct {
	Tag                string
	Name               string
	MapPosition        uint32
	TownhallLevel      uint32
	OpponentAttacks    uint32
	BestOpponentAttack ClanWarAttack
	Attacks            []ClanWarAttack
}

// ClanWarAttack 部落战进攻
type ClanWarAttack struct {
	Order                 uint32
	AttackerTag           string
	DefenderTag           string
	Stars                 uint32
	DestructionPercentage uint32
	Duration              uint32
}
