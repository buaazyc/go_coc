package parser

import "encoding/json"

func LeagueGroup(info string) (*ClanWarLeagueGroup, error) {
	var leagueGroup *ClanWarLeagueGroup
	err := json.Unmarshal([]byte(info), &leagueGroup)
	if err != nil {
		return nil, err
	}
	return leagueGroup, nil
}

type LeagueGroupRsp struct {
	League string     `json:"league"`
	Season string     `json:"season"`
	Clans  [][]string `json:"clans"`
}

type ClanWarLeagueGroup struct {
	Tag    string               `json:"tag"`
	State  string               `json:"state"`
	Season string               `json:"season"`
	Clans  []ClanWarLeagueClan  `json:"clans"`
	Rounds []ClanWarLeagueRound `json:"rounds"`
}

type ClanWarLeagueClan struct {
	Tag       string                    `json:"tag"`
	ClanLevel uint32                    `json:"clanLevel"`
	Name      string                    `json:"name"`
	Members   []ClanWarLeagueClanMember `json:"members"`
	BadgeUrls BadgeUrl                  `json:"badgeUrls"`
}

type ClanWarLeagueRound struct {
	WarTags []string `json:"warTags"`
}

type ClanWarLeagueClanMember struct {
	Tag           string `json:"tag"`
	TownHallLevel uint32 `json:"townHallLevel"`
	Name          string `json:"name"`
}
