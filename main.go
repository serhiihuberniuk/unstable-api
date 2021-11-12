package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/serhiihuberniuk/unstable-api/config"
	"github.com/serhiihuberniuk/unstable-api/handler"
	"github.com/serhiihuberniuk/unstable-api/storage/cache/custom"
	"github.com/serhiihuberniuk/unstable-api/storage/decorator/cache"
	"github.com/serhiihuberniuk/unstable-api/storage/decorator/error_generator"
	"github.com/serhiihuberniuk/unstable-api/storage/remote/opendota"
)

func main() {

	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("error while reading configs: %v", err)
	}

	ctx := context.Background()
	f := opendota.New()
	c := custom.NewCustomStorage(ctx, time.Duration(cfg.CacheExpirationTime), time.Duration(cfg.CacheCleanupInterval))

	cacheDecorator := cache.NewCacheDecorator(f, c)
	errorDecorator := error_generator.NewErrorGeneratorDecorator(cacheDecorator)
	h := handler.New(errorDecorator)

	srv := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: h.ApiRouter(),
	}

	errs := make(chan error, 0)

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			errs <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errs:
		log.Fatalf("error occured while running server: %v", err)
	case <-quit:
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	if err = srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		cancel()
		log.Fatalf("error while shutting down the server: %v", err)
	}

	cancel()
}
