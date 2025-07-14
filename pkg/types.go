package pkg

import (
	"fmt"
	"strings"
)

// Color constants for better maintainability
const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorBlue    = "\033[34m"
	ColorGreen   = "\033[32m"
	ColorWhite   = "\033[37m"
	ColorYellow  = "\033[33m"
	ColorMagenta = "\033[95m"
	ColorCyan    = "\033[36m"
	ColorPurple  = "\033[35m"
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

func (p *PvPEntry) Render() {
	var parts []string
	parts = append(parts, fmt.Sprintf("%-6d", p.Rating))
	parts = append(parts, fmt.Sprintf("%-22s", p.FactionNameColor()))

	parts = append(parts, fmt.Sprintf("%-15s", p.SeasonMatchStatistics.Winrate()))

	if p.CharacterResponse != nil {
		parts = append(parts, fmt.Sprintf("  %-15s %-18s",
			p.CharacterResponse.ClassSpecColor(),
			p.Race.Name))
	}
	
	fmt.Println(strings.Join(parts, ""))
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

func (p *PvPEntry) FactionNameColor() string {
	switch p.Faction.Type {
	case "HORDE":
		return ColorRed + p.Character.Name + ColorReset
	case "ALLIANCE":
		return ColorBlue + p.Character.Name + ColorReset
	default:
		return p.Character.Name
	}
}

func (c *CharacterResponse) ClassSpecColor() string {
	color, exists := classColors[c.CharacterClass.Id]
	if !exists {
		return fmt.Sprintf("%-25s", c.ActiveSpec.Name+" "+c.CharacterClass.Name)
	}

	classSpec := c.ActiveSpec.Name + " " + c.CharacterClass.Name
	return fmt.Sprintf("%s%-20s%s", color, classSpec, ColorReset)
}

func (s *SeasonMatchStatistics) Winrate() string {
	winLoss := fmt.Sprintf("%s%3d%s/%s%-3d%s",
		ColorGreen, s.Won, ColorReset,
		ColorRed, s.Lost, ColorReset)

	percentage := (float64(s.Won) / float64(s.Played)) * 100
	percentStr := fmt.Sprintf("%.2f%%", percentage)

	return fmt.Sprintf(" %-3s %s", winLoss, percentStr)
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

var classColors = map[int]string{
	1:  "\033[38;5;130m", // Warrior - Brown
	2:  ColorMagenta,     // Paladin - Magenta
	3:  ColorGreen,       // Hunter - Green
	4:  ColorYellow,      // Rogue - Yellow
	5:  ColorWhite,       // Priest - White
	6:  ColorRed,         // Death Knight - Red
	7:  ColorBlue,        // Shaman - Blue
	8:  ColorCyan,        // Mage - Cyan
	9:  "\033[38;5;129m", // Warlock - Purple
	10: ColorGreen,       // Monk - Green
	11: "\033[38;5;208m", // Druid - Orange
	12: ColorPurple,      // Demon Hunter - Purple
	13: "\033[38;5;82m",  // Evoker - Light Green
}
