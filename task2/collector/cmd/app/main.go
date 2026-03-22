package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"github.com/Friend-zva/golang-course-task2/collector/config"
	"github.com/Friend-zva/golang-course-task2/collector/internal/adapter/github"
	grpcH "github.com/Friend-zva/golang-course-task2/collector/internal/controller/grpc"
	"github.com/Friend-zva/golang-course-task2/collector/internal/usecase"
	pb "github.com/Friend-zva/golang-course-task2/proto/pkg/api/v1"
)

func main() {
	cfg := config.MustLoad("config/local.yaml")

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
	pb.RegisterInfoRepoServiceServer(grpcServer, handler)

	log.Printf("starting server on %s", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
