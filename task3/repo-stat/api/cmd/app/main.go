package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/Friend-zva/golang-course-task3/repo-stat/api/config"
	"github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/controller/http"
	"github.com/Friend-zva/golang-course-task3/repo-stat/platform/httpserver"
	"github.com/Friend-zva/golang-course-task3/repo-stat/platform/logger"
)

func run(ctx context.Context) error {
	// config
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "server configuration file")
	flag.Parse()

	cfg := config.MustLoad(configPath)

	// logger

	log := logger.MustMakeLogger(cfg.Logger.LogLevel)

	log.Info("starting server...")
	log.Debug("debug messages are enabled")

	// handler
	handler, err := http.NewHandler(ctx, log, cfg)
	if err != nil {
		log.Error("Error creating handler", "error", err)
		return err
	}

	// server
	srv := httpserver.New(cfg.HTTP, handler)
	if err := srv.Run(ctx); err != nil {
		return fmt.Errorf("run http server: %w", err)
	}
	return nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	if err := run(ctx); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		if err != nil {
			fmt.Printf("launching server error: %s\n", err)
		}
		cancel()
		os.Exit(1)
	}
	cancel()
}
