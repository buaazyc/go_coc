package parser

import "encoding/json"

// PlayerInfo 解析玩家信息
func PlayerInfo(info string) (*Player, error) {
	var playerInfo *Player
	err := json.Unmarshal([]byte(info), &playerInfo)
	if err != nil {
		return nil, err
	}
	return playerInfo, nil
}

type Player struct {
	Tag           string `json:"tag"`
	Name          string `json:"name"`
	Role          string `json:"role"`
	WarPreference string `json:"warPreference"`

	ExpLevel            uint32 `json:"expLevel"`
	Trophies            uint32 `json:"trophies"`
	BestTrophies        uint32 `json:"bestTrophies"`
	Donations           uint32 `json:"donations"`
	DonationsReceived   uint32 `json:"donationsReceived"`
	WarStars            uint32 `json:"warStars"`
	AttackWins          uint32 `json:"attackWins"`
	DefenseWins         uint32 `json:"defenseWins"`
	TownHallLevel       uint32 `json:"townHallLevel"`
	TownHallWeaponLevel uint32 `json:"townHallWeaponLevel"`

	BuilderHallLevel     uint32 `json:"builderHallLevel"`
	VersusTrophies       uint32 `json:"versusTrophies"`
	BestVersusTrophies   uint32 `json:"bestVersusTrophies"`
	VersusBattleWins     uint32 `json:"versusBattleWins"`
	VersusBattleWinCount uint32 `json:"versusBattleWinCount"`

	Clan             *PlayerClan                  `json:"clan"`
	League           *League                      `json:"league"`
	LegendStatistics *PlayerLegendStatistics      `json:"legendStatistics"`
	Spells           []*PlayerItemLevel           `json:"spells"`
	Troops           []*PlayerItemLevel           `json:"troops"`
	Heroes           []*PlayerItemLevel           `json:"heroes"`
	Labels           []*Label                     `json:"labels"`
	Achievements     []*PlayerAchievementProgress `json:"achievements"`
}

type PlayerClan struct {
	Tag       string   `json:"tag"`
	ClanLevel uint32   `json:"clanLevel"`
	Name      string   `json:"name"`
	BadgeUrls BadgeUrl `json:"badgeUrls"`
}

type PlayerItemLevel struct {
	Level              uint32 `json:"level"`
	Name               string `json:"name"`
	MaxLevel           uint32 `json:"maxLevel"`
	Village            string `json:"village"`
	SuperTroopIsActive bool   `json:"superTroopIsActive"`
}

type PlayerLegendStatistics struct {
	LegendTrophies       uint32                             `json:"legendTrophies"`
	PreviousSeason       LegendLeagueTournamentSeasonResult `json:"previousSeason"`
	PreviousVersusSeason LegendLeagueTournamentSeasonResult `json:"previousVersusSeason"`
	BestVersusSeason     LegendLeagueTournamentSeasonResult `json:"bestVersusSeason"`
	CurrentSeason        LegendLeagueTournamentSeasonResult `json:"currentSeason"`
	BestSeason           LegendLeagueTournamentSeasonResult `json:"bestSeason"`
}

type LegendLeagueTournamentSeasonResult struct {
	ID       string `json:"id"`
	Rank     uint32 `json:"rank"`
	Trophies uint32 `json:"trophies"`
}

type PlayerAchievementProgress struct {
	Stars          uint32 `json:"stars"`
	Value          uint32 `json:"value"`
	Target         uint32 `json:"target"`
	Name           string `json:"name"`
	Info           string `json:"info"`
	CompletionInfo string `json:"completionInfo"`
	Village        string `json:"village"`
}
