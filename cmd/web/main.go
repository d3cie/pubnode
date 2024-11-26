package main

import (
	"log/slog"
	"os"

	"github.com/d3cie/pubnode/internal/config"
	"github.com/d3cie/pubnode/internal/infra/db"
	"github.com/d3cie/pubnode/pkg/utils"
	"github.com/lmittmann/tint"
)

type app struct {
	db     *db.DB
	logger *slog.Logger
	cfg    *config.Config
}

func main() {
	config.Init()
	cfg := config.Get()

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: utils.Ternary(cfg.Dev, slog.LevelDebug, slog.LevelInfo)}))
	db, err := db.New()
	if err != nil {
		logger.Error("failed to initialize database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	a := &app{
		db:     db,
		logger: logger,
		cfg:    cfg,
	}
	if err := a.startServer(); err != nil {
		logger.Error("failed to start server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
