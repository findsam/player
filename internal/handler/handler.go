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
	Client      *resty.Client
}

func NewHandler(t *pkg.Token) *Handler {
	return &Handler{
		AccessToken: t.AccessToken,
		Client:      t.Client,
	}
}
func (h *Handler) GetLeaderboard() (*pkg.PvPLeaderboard, error) {
	result := &pkg.PvPLeaderboard{}
	stop := pkg.StartSpinner()

	resp, err := h.Client.R().
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
