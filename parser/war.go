package parser

import (
	"encoding/json"
)

// CurrentWar 解析部落战信息
func CurrentWar(info string) (*ClanWar, error) {
	var warInfo *ClanWar
	err := json.Unmarshal([]byte(info), &warInfo)
	if err != nil {
		return nil, err
	}
	return warInfo, nil
}

// ClanWar 当前部落战信息
type ClanWar struct {
	Clan                 *WarClan `json:"clan"`
	Opponent             *WarClan `json:"opponent"`
	TeamSize             uint32   `json:"teamSize"`
	StartTime            string   `json:"startTime"`
	EndTime              string   `json:"endTime"`
	State                string   `json:"state"`
	PreparationStartTime string   `json:"preparationStartTime"`
}

// WarClan 参战部落信息
type WarClan struct {
	DestructionPercentage float64          `json:"destructionPercentage"`
	Tag                   string           `json:"tag"`
	Name                  string           `json:"name"`
	BadgeUrls             *BadgeUrl        `json:"badgeUrls"`
	ClanLevel             uint32           `json:"clanLevel"`
	Attacks               uint32           `json:"attacks"`
	Stars                 uint32           `json:"stars"`
	ExpEarned             uint32           `json:"expEarned"`
	Members               []*ClanWarMember `json:"members"`
}

// BadgeUrl 部落图标
type BadgeUrl struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Samll  string `json:"samll"`
}

// ClanWarMember 参战成员信息
type ClanWarMember struct {
	Tag                string           `json:"tag"`
	Name               string           `json:"name"`
	MapPosition        uint32           `json:"mapPosition"`
	TownhallLevel      uint32           `json:"townhallLevel"`
	OpponentAttacks    uint32           `json:"opponentAttacks"`
	BestOpponentAttack *ClanWarAttack   `json:"bestOpponentAttack"`
	Attacks            []*ClanWarAttack `json:"attacks"`
}

// ClanWarAttack 部落战进攻
type ClanWarAttack struct {
	Order                 uint32 `json:"order"`
	AttackerTag           string `json:"attackerTag"`
	DefenderTag           string `json:"defenderTag"`
	Stars                 uint32 `json:"stars"`
	DestructionPercentage uint32 `json:"destructionPercentage"`
	Duration              uint32 `json:"duration"`
}
