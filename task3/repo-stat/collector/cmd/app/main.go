package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	config "github.com/Friend-zva/golang-course-task3/repo-stat/collector/config"
	github "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/adapter/github"
	grpccontroller "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/controller/grpc"
	usecase "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/usecase"
	grpcserver "github.com/Friend-zva/golang-course-task3/repo-stat/platform/grpcserver"
	logger "github.com/Friend-zva/golang-course-task3/repo-stat/platform/logger"
	collectorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/collector"
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

	// github client
	httpClient := &http.Client{
		Timeout: cfg.GRPC.Timeout,
	}
	clientGH := github.NewClient(httpClient, log)

	// handler
	getInfoRepo := usecase.NewGetInfoRepo(clientGH)
	handler := grpccontroller.NewInfoRepoHandler(log, getInfoRepo)

	// server
	server, err := grpcserver.New(cfg.GRPC.Address)
	if err != nil {
		return fmt.Errorf("cannot create grpc server: %w", err)
	}

	collectorpb.RegisterCollectorServer(server.GRPC(), handler)

	if err := server.Run(ctx); err != nil {
		return fmt.Errorf("cannot run grpc server: %w", err)
	}

	return nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	if err := run(ctx); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		if err != nil {
			fmt.Printf("cannot launch collector server: %s\n", err)
		}
		cancel()
		os.Exit(1)
	}

	cancel()
}
