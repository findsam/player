package render

import (
	"fmt"
	"sync"
	"time"

	"github.com/findsam/tbot/internal/handler"
	"github.com/findsam/tbot/internal/repo"
	"github.com/findsam/tbot/pkg"
)

type Render struct {
	Handler *handler.Handler
	Repo    *repo.Repo
}

func NewRender(h *handler.Handler, r *repo.Repo) *Render {
	return &Render{Handler: h, Repo: r}
}

func (r *Render) List() error {
	start := time.Now()
	res, err := r.Handler.GetLeaderboard()
	if err != nil {
		return fmt.Errorf("failed to get leaderboard: %w", err)
	}

	const batchSize = 10
	entries := res.Entries[:100]
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
	}

	for _, player := range entries {
		fmt.Printf("rendering: %s\n", player.Character.Name)
		r.Repo.SavePlayer(&player)
	}

	close(eChan)
	for err := range eChan {
		fmt.Println(err)
	}

	fmt.Println(time.Since(start))
	return nil
}

func (r *Render) GetList() error {
	r.Repo.GetPlayers()
	return nil
}