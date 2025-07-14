package cmd

import (
	"fmt"

	"github.com/findsam/tbot/internal/handler"
	"github.com/findsam/tbot/internal/render"
	"github.com/findsam/tbot/pkg"
)

func Execute() error {
	token := pkg.NewToken()
	if err := token.Get(); err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	h := handler.NewHandler(token)
	renderer := render.NewRender(h)
	
	if err := renderer.List(); err != nil {
		return fmt.Errorf("failed to list: %w", err)
	}
	
	return nil
}