package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kikplate-plates/go-clean-architecture-starter/internal/app"
	"github.com/kikplate-plates/go-clean-architecture-starter/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("config_error", slog.Any("err", err))
		os.Exit(1)
	}
	level := slog.LevelInfo
	if cfg.LogLevel == "debug" {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	bootstrapCtx, bootCancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer bootCancel()
	a, err := app.New(bootstrapCtx, cfg)
	if err != nil {
		slog.Error("bootstrap_error", slog.Any("err", err))
		os.Exit(1)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	if err := a.Run(ctx, shutdownCtx); err != nil {
		slog.Error("exit_error", slog.Any("err", err))
		os.Exit(1)
	}
}
