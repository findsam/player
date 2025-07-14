package render

import (
	"fmt"
	"sync"

	"github.com/findsam/tbot/internal/handler"
	"github.com/findsam/tbot/pkg"
)

type Render struct {
	Handler *handler.Handler
}

func NewRender(h *handler.Handler) *Render {
	return &Render{Handler: h}
}

func (r *Render) List() error {
	res, err := r.Handler.GetLeaderboard()
	if err != nil {
		return fmt.Errorf("failed to get leaderboard: %w", err)
	}

	entries := res.Entries
	if len(entries) > 25 {
		entries = entries[:25]
	}

	return r.upsertPlayers(entries)
}

func (r *Render) upsertPlayers(entries []pkg.PvPEntry) error {
	const batchSize = 5

	for start := 0; start < len(entries); start += batchSize {
		end := min(start + batchSize, len(entries))

		batch := entries[start:end]
		var wg sync.WaitGroup
		updatedBatch := make([]pkg.PvPEntry, len(batch))

		for i, player := range batch {
			wg.Add(1)
			go func(i int, p pkg.PvPEntry) {
				defer wg.Done()
				resp, err := r.Handler.GetPlayer(p.Character.Realm.Slug, p.Character.Name)

				if err != nil { 
					fmt.Println(err)
				}
				
				p.CharacterResponse = resp
				updatedBatch[i] = p
			}(i, player)
		}

		wg.Wait()
		copy(entries[start:end], updatedBatch)
		r.renderList(entries)
	}

	return nil
}

func (r *Render) renderList(entries []pkg.PvPEntry) {
	fmt.Print("\033[2J\033[H")
	for _, player := range entries {
		player.Render()
	}
}
