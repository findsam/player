package cmd

import (
	"fmt"
	"time"

	"github.com/findsam/tbot/internal/handler"
	"github.com/findsam/tbot/internal/render"
	"github.com/findsam/tbot/pkg"
	"resty.dev/v3"
)

func Execute() error {
	c := resty.New().SetTimeout(30 * time.Second)
	t := pkg.NewToken(c)
	if err := t.Get(); err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}
	
	h := handler.NewHandler(t)
	r := render.NewRender(h)
	
	if err := r.List(); err != nil {
		return fmt.Errorf("failed to list: %w", err)
	}
	
	return nil
}