package parser

import "encoding/json"

func ClanInfo(info string) (*Clan, error) {
	var clan *Clan
	err := json.Unmarshal([]byte(info), &clan)
	if err != nil {
		return nil, err
	}
	return clan, nil
}

type Clan struct {
	Tag                    string        `json:"tag"`
	ClanLevel              uint32        `json:"clanLevel"`
	WarWinStreak           uint32        `json:"warWinStreak"`
	WarWins                uint32        `json:"warWins"`
	WarTies                uint32        `json:"warTies"`
	WarLosses              uint32        `json:"warLosses"`
	ClanPoints             uint32        `json:"clanPoints"`
	Name                   string        `json:"name"`
	Members                uint32        `json:"members"`
	Type                   string        `json:"type"`
	Description            string        `json:"description"`
	ClanVersusPoints       uint32        `json:"clanVersusPoints"`
	RequiredTrophies       uint32        `json:"requiredTrophies"`
	RequiredVersusTrophies uint32        `json:"requiredVersusTrophies"`
	RequiredTownhallLevel  uint32        `json:"requiredTownhallLevel"`
	IsWarLogPublic         bool          `json:"isWarLogPublic"`
	WarFrequency           string        `json:"warFrequency"`
	ChatLanguage           *Language     `json:"chatLanguage"`
	WarLeague              *WarLeague    `json:"warLeague"`
	MemberList             []*ClanMember `json:"memberList"`
	Labels                 []*Label      `json:"labels"`
	Location               *Location     `json:"location"`
	BadgeUrls              *BadgeUrl     `json:"badgeUrls"`
}

type WarLeague struct {
	Name string `json:"name"`
	ID   uint32 `json:"id"`
}

type Language struct {
	Name         string `json:"name"`
	ID           uint32 `json:"id"`
	LanguageCode string `json:"languageCode"`
}

type Location struct {
	LocalizedName string `json:"localizedName"`
	ID            uint32 `json:"id"`
	Name          string `json:"name"`
	IsCountry     bool   `json:"isCountry"`
	CountryCode   string `json:"countryCode"`
}

type Label struct {
	Name     string    `json:"name"`
	ID       uint32    `json:"id"`
	IconUrls *IconUrls `json:"iconUrls"`
}

type ClanMember struct {
	Tag               string  `json:"tag"`
	Name              string  `json:"name"`
	Role              string  `json:"role"`
	ExpLevel          uint32  `json:"expLevel"`
	ClanRank          uint32  `json:"clanRank"`
	PreviousClanRank  uint32  `json:"previousClanRank"`
	Donations         uint32  `json:"donations"`
	DonationsReceived uint32  `json:"donationsReceived"`
	Trophies          uint32  `json:"trophies"`
	VersusTrophies    uint32  `json:"versusTrophies"`
	League            *League `json:"league"`
}

type League struct {
	Name     string    `json:"name"`
	ID       uint32    `json:"id"`
	IconUrls *IconUrls `json:"iconUrls"`
}

type IconUrls struct {
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
	Medium string `json:"medium"`
}
