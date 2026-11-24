package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThanhThien334/MangaHub/internal/config"
	"github.com/ThanhThien334/MangaHub/internal/httpapi"
	"github.com/ThanhThien334/MangaHub/internal/logging"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	logger := logging.New("api", cfg.LogLevel)
	if cfg.UsingDefaultSecret() {
		logger.Warn("using default JWT secret")
	}

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpapi.NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErr:
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("api server stopped", "error", err)
			return err
		}
		return nil
	case <-ctx.Done():
		// Let in-flight requests finish before stopping the server.
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("api shutdown failed", "error", err)
			return err
		}
		if err := <-serverErr; err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("api server stopped", "error", err)
			return err
		}
		logger.Info("api server stopped")
		return nil
	}
}
