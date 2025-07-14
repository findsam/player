package pkg

import (
	"fmt"
	"time"

	"resty.dev/v3"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) Get() error {
	client := resty.New().
		SetTimeout(30 * time.Second)

	resp, err := client.R().
		SetFormData(map[string]string{
			"grant_type": "client_credentials",
		}).
		SetBasicAuth(Envs.BLIZZARD_CLIENT_ID, Envs.BLIZZARD_CLIENT_SECRET).
		SetResult(t).
		Post("https://oauth.battle.net/token")

	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode(), resp.String())
	}
	return nil
}
