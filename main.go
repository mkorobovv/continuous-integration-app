package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/mkorobovv/continuous-integration-app/internal/app/config"
	"github.com/mkorobovv/continuous-integration-app/internal/app/httpadapter"
	"github.com/mkorobovv/continuous-integration-app/internal/app/infrastructure/redis"
	"github.com/mkorobovv/continuous-integration-app/internal/app/repositories/globalincrementrepository"
	"github.com/mkorobovv/continuous-integration-app/internal/app/repositories/shortenrepository"
	"github.com/mkorobovv/continuous-integration-app/internal/app/services/urlshortenerservice"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	shortenDB := redis.New(cfg.Databases.Shorten)

	shortenRepository := shortenrepository.New(shortenDB)
	globalIncrementRepository := globalincrementrepository.New(shortenDB)

	urlShortenService := urlshortenerservice.New(shortenRepository, globalIncrementRepository)

	httpAdapter := httpadapter.New(urlShortenService)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		slog.Info("Starting server", slog.String("port", cfg.HTTP.Port))

		return httpAdapter.Start(gCtx)
	})

	err = g.Wait()
	if err != nil {
		if !errors.Is(err, context.Canceled) {
			logger.Error("application execution failed", "error", err)
			os.Exit(1)
		}
	}
}
