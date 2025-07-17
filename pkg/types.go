package pkg

import (
	"fmt"
)

type PvPLeaderboard struct {
	Entries []PvPEntry `json:"entries"`
}

type PvPEntry struct {
	Character             Character             `json:"character"`
	Faction               Faction               `json:"faction"`
	Rank                  int                   `json:"rank"`
	Rating                int                   `json:"rating"`
	SeasonMatchStatistics SeasonMatchStatistics `json:"season_match_statistics"`
	*CharacterResponse
}

type Character struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Realm Realm  `json:"realm"`
}

type Realm struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}

type SeasonMatchStatistics struct {
	Lost   int `json:"lost"`
	Played int `json:"played"`
	Won    int `json:"won"`
}

type Faction struct {
	Type string `json:"type"`
}

func (s *SeasonMatchStatistics) Winrate() (string, string) {
	winLoss := fmt.Sprintf("%v/%v", s.Won, s.Lost)
	percentage := (float64(s.Won) / float64(s.Played)) * 100
	percentStr := fmt.Sprintf("%.2f%%", percentage)
	return winLoss, percentStr
}

type KeyReference struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type CharacterResponse struct {
	Gender         KeyReference `json:"gender"`
	Race           KeyReference `json:"race"`
	CharacterClass KeyReference `json:"character_class"`
	ActiveSpec     KeyReference `json:"active_spec"`
}
