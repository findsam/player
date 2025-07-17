package repo

import (
	"fmt"

	"github.com/findsam/tbot/pkg"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type Leaderboard struct {
	Rating        int
	Race          string
	Name          string
	Class         string
	Realm         string
	Winloss       string
	Statistics    string
	Specilisation string
}

func (r *Repo) SavePlayer(data *pkg.PvPEntry) error {
	winLoss, stats := data.SeasonMatchStatistics.Winrate()
	input := &Leaderboard{
		Name:   data.Character.Name,
		Rating: data.Rating,
		Winloss:       winLoss,
		Statistics:    stats,
		Realm: data.Character.Realm.Slug,
	}
	
	if data.CharacterResponse != nil {
		input.Specilisation = data.ActiveSpec.Name
		input.Class = data.CharacterClass.Name
		input.Race = data.Race.Name
	}

	result := r.db.Create(input)
	if result.Error != nil {
		return fmt.Errorf("err")
	}
	return nil
}

func (r *Repo) GetPlayers() (*[]Leaderboard, error) {
	var players []Leaderboard
	r.db.Order("rating desc").Find(&players)
	for _, player := range players {
		fmt.Println(player)
	}
	return nil, nil 
}