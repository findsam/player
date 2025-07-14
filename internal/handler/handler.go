package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/findsam/tbot/pkg"
	"resty.dev/v3"
)

type Handler struct {
	AccessToken string
}

func NewHandler(t *pkg.Token) *Handler {
	return &Handler{
		AccessToken: t.AccessToken,
	}
}
func StartSpinner() (stop func()) {
	done := make(chan struct{})
	go func() {
		frames := []string{"|", "/", "-", "\\"}
		i := 0
		ticker := time.NewTicker(250 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Printf("\r%s", frames[i%len(frames)])
				i++
			}
		}
	}()
	return func() {
		close(done)
		fmt.Print("\r\033[K")
	}
}

func (h *Handler) GetLeaderboard() (*pkg.PvPLeaderboard, error) {
	client := resty.New().SetTimeout(30 * time.Second)
	result := &pkg.PvPLeaderboard{}
	stop := StartSpinner()

	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", h.AccessToken)).
		SetResult(result).
		Get("https://us.api.blizzard.com/data/wow/pvp-season/39/pvp-leaderboard/3v3?namespace=dynamic-us&locale=en_US")

	stop()

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode(), resp.String())
	}

	return result, nil
}

func (h *Handler) GetPlayer(slug, name string) (*pkg.CharacterResponse, error) {
	client := resty.New().SetTimeout(30 * time.Second)
	result := &pkg.CharacterResponse{}

	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", h.AccessToken)).
		SetResult(result).
		Get(fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s?namespace=profile-us&locale=en_US", strings.ToLower(slug), strings.ToLower(name)))

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode(), resp.String())
	}

	return result, nil
}
