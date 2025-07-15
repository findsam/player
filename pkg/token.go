package pkg

import (
	"fmt"

	"resty.dev/v3"
)

type Token struct {
	AccessToken string `json:"access_token"`
	Client      *resty.Client
}

func NewToken(c *resty.Client) *Token {
	return &Token{
		Client: c,
	}
}

func (t *Token) Get() error {
	resp, err := t.Client.R().
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
