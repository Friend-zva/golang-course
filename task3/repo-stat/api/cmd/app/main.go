package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	config "github.com/Friend-zva/golang-course-task3/repo-stat/api/config"
	processor "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/adapter/processor"
	subscriber "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/adapter/subscriber"
	httpcontroller "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/controller/http"
	usecase "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/usecase"
	httpserver "github.com/Friend-zva/golang-course-task3/repo-stat/platform/httpserver"
	logger "github.com/Friend-zva/golang-course-task3/repo-stat/platform/logger"
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

	// subscriber client
	clientSubs, err := subscriber.NewClient(cfg.Services.Subscriber, log)
	if err != nil {
		return fmt.Errorf("cannot init subscriber client: %w", err)
	}
	defer func() {
		if err := clientSubs.Close(); err != nil {
			log.Warn("cannot close subscriber connection", "error", err)
		}
	}()

	// processor client
	clientProc, err := processor.NewClient(cfg.Services.Processor, log)
	if err != nil {
		return fmt.Errorf("cannot init processor client: %w", err)
	}
	defer func() {
		if err := clientProc.Close(); err != nil {
			log.Warn("cannot close processor connection", "error", err)
		}
	}()

	// handler
	pingSubs := usecase.NewPing(clientSubs)
	pingProc := usecase.NewPing(clientProc)
	getInfoRepo := usecase.NewGetInfoRepo(clientProc)
	handler := httpcontroller.NewRouter(log, pingSubs, pingProc, getInfoRepo)

	// server
	server := httpserver.New(cfg.HTTP, handler)
	if err := server.Run(ctx); err != nil {
		return fmt.Errorf("cannot run http server: %w", err)
	}

	return nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	if err := run(ctx); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		if err != nil {
			fmt.Printf("cannot launch api server: %s\n", err)
		}
		cancel()
		os.Exit(1)
	}

	cancel()
}
