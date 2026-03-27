package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	grpcH "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/controller/grpc"
	collectorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/collector"

	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/config"
	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/adapter/github"
	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/usecase"
)

func main() {
	cfg := config.MustLoad("config/config.yaml")

	listener, err := net.Listen("tcp", cfg.CollectorServer.Address)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	httpClient := &http.Client{
		Timeout: cfg.CollectorServer.Timeout,
	}
	clientGH := github.NewGitHubAPI(httpClient)

	info := usecase.NewInfoRepo(clientGH)

	handler := grpcH.NewHandler(info)

	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(cfg.CollectorServer.IdleTimeout),
	)
	collectorpb.RegisterCollectorServer(grpcServer, handler)

	go func() {
		log.Printf("Starting server on %s...", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Printf("failed to serve: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Collector gRPC server...")

	grpcServer.GracefulStop()
	log.Println("Collector server exited properly")
}
