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

	const batchSize = 5
	entries := res.Entries[:20]
	eChan := make(chan error, len(entries))

	for start := 0; start < len(entries); start += batchSize {
		end := min(start+batchSize, len(entries))
		batch := entries[start:end]
		var wg sync.WaitGroup
		list := make([]pkg.PvPEntry, len(batch))

		for i, player := range batch {
			wg.Add(1)
			go func(i int, p pkg.PvPEntry) {
				defer wg.Done()
				resp, err := r.Handler.GetPlayer(p.Character.Realm.Slug, p.Character.Name)
				if err != nil {
					eChan <- fmt.Errorf("error fetching %s's dynamic details", p.Character.Name)
				}
				p.CharacterResponse = resp
				list[i] = p
			}(i, player)
		}

		wg.Wait()
		copy(entries[start:end], list)

		fmt.Print("\033[2J\033[H")
		for _, player := range entries {
			player.Render()
		}
	}

	close(eChan)
	for err := range eChan {
		fmt.Println(err)
	}

	return nil
}
