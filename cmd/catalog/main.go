package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThanhThien334/MangaHub/internal/config"
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
	logger := logging.New("catalog", cfg.LogLevel)

	listener, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		logger.Error("catalog listener failed", "addr", cfg.GRPCAddr, "error", err)
		return err
	}
	logger.Info("catalog listening", "addr", listener.Addr().String())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	if err := listener.Close(); err != nil {
		logger.Error("catalog listener close failed", "error", err)
		return err
	}
	logger.Info("catalog stopped")
	return nil
}
