package main

import (
	"context"
	"home-hap/internal/config"
	"home-hap/pkg/discovery"
	"home-hap/pkg/logging"
	"home-hap/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create context
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Load logger
	logging.Init(cfg.Log)
	logger := logging.GetLogger()
	logger.Info("Starting app")

	// Discovery devices
	disco := discovery.New(ctx, cfg)
	disco.LocalDevices()
	// --

	hap := server.NewHomeKit(ctx, cfg.HomeKit, disco)
	hap.Start()

	_ = <-c // This blocks the main thread until an interrupt is received
	logger.Info("Gracefully shutting down...")
	cancel()
	logger.Sync()
}
