package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/findsam/tbot/internal/handler"
	"github.com/findsam/tbot/internal/render"
	"github.com/findsam/tbot/pkg"
	"resty.dev/v3"
)

func Execute() error {
	db, err := pkg.NewDB(pkg.Envs.DB_USER, pkg.Envs.DB_PWD, pkg.Envs.DB_NAME).Start()

	if err != nil { 
		return fmt.Errorf("failed to start DB %w", err)
	}

	db.Ping(context.Background());

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
