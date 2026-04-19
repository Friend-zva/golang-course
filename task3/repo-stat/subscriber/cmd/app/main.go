package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	grpcserver "github.com/Friend-zva/golang-course-task3/repo-stat/platform/grpcserver"
	logger "github.com/Friend-zva/golang-course-task3/repo-stat/platform/logger"
	subscriberpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/subscriber"
	"github.com/Friend-zva/golang-course-task3/repo-stat/subscriber/config"
	grpccontroller "github.com/Friend-zva/golang-course-task3/repo-stat/subscriber/internal/controller/grpc"
	usecase "github.com/Friend-zva/golang-course-task3/repo-stat/subscriber/internal/usecase"
)

func run(ctx context.Context) error {
	// config
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "server configuration file")
	flag.Parse()
	cfg := config.MustLoad(configPath)

	// logger
	log := logger.MustMakeLogger(cfg.Logger.LogLevel)
	log.Info("starting subscriber server...")
	log.Debug("debug messages are enabled")

	// handler
	pingUseCase := usecase.NewPing()
	pingServer := grpccontroller.NewServer(log, pingUseCase)

	// server
	srv, err := grpcserver.New(cfg.GRPC.Address)
	if err != nil {
		return fmt.Errorf("cannot create grpc server: %w", err)
	}

	subscriberpb.RegisterSubscriberServer(srv.GRPC(), pingServer)

	if err := srv.Run(ctx); err != nil {
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
