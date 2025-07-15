package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/findsam/tbot/internal/handler"
	"github.com/findsam/tbot/internal/render"
	"github.com/findsam/tbot/pkg"
	"resty.dev/v3"
)

func Execute() error {
	db := pkg.NewDB(pkg.Envs.DB_USER, pkg.Envs.DB_PWD, pkg.Envs.DB_NAME)

	conn, err := db.Start()
	if err != nil {
		return fmt.Errorf("failed to start DB: %w", err)
	}
	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Printf("failed to close DB connection: %v", err)
		}
	}()

	if err := db.Migrate(); err != nil {
		return fmt.Errorf("failed to migrate DB: %w", err)
	}

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
