package cmd

import (
	"fmt"
	"time"

	"github.com/findsam/tbot/internal/handler"
	"github.com/findsam/tbot/internal/render"
	"github.com/findsam/tbot/internal/repo"
	"github.com/findsam/tbot/pkg"
	"resty.dev/v3"
)

func Execute() error {
	db, err := repo.NewDB(pkg.Envs.DB_USER, pkg.Envs.DB_PWD, pkg.Envs.DB_NAME).Start()

	if err != nil {
		return fmt.Errorf("db err")
	}

	r := repo.NewRepo(db)
	c := resty.New().SetTimeout(30 * time.Second)

	t := pkg.NewToken(c)
	if err := t.Get(); err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	h := handler.NewHandler(t)
	rn := render.NewRender(h, r)
	if err := rn.List(); err != nil {
		return fmt.Errorf("failed to list: %w", err)
	}

	rn.GetList()
	return nil
}
