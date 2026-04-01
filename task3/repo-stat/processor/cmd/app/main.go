package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	grpcserver "github.com/Friend-zva/golang-course-task3/repo-stat/platform/grpcserver"
	logger "github.com/Friend-zva/golang-course-task3/repo-stat/platform/logger"
	config "github.com/Friend-zva/golang-course-task3/repo-stat/processor/config"
	collector "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/adapter/collector"
	grpcprocessor "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/controller/grpc"
	usecase "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/usecase"
	processorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/processor"
)

func run(ctx context.Context) error {
	// config
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "server configuration file")
	flag.Parse()
	cfg := config.MustLoad(configPath)

	// logger
	log := logger.MustMakeLogger(cfg.Logger.LogLevel)
	log.Info("starting processor server...")
	log.Debug("debug messages are enabled")

	// collector client
	clientColl, err := collector.NewClient(cfg.Services.Collector, log)
	if err != nil {
		return fmt.Errorf("cannot init collector client: %w", err)
	}
	defer func() {
		if err := clientColl.Close(); err != nil {
			log.Warn("cannot close collector connection", "error", err)
		}
	}()

	// handler
	ping := usecase.NewPing()
	getInfoRepo := usecase.NewGetInfoRepo(clientColl)
	handler := grpcprocessor.NewHandler(log, ping, getInfoRepo)

	// server
	server, err := grpcserver.New(cfg.GRPC.Address)
	if err != nil {
		return fmt.Errorf("cannot create grpc server: %w", err)
	}

	processorpb.RegisterProcessorServer(server.GRPC(), handler)

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
			fmt.Printf("cannot launch processor server: %v\n", err)
		}
		cancel()
		os.Exit(1)
	}

	cancel()
}
