package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"github.com/Friend-zva/golang-course-task2/collector/internal/adapter/github"
	grpcS "github.com/Friend-zva/golang-course-task2/collector/internal/controller/grpc"
	"github.com/Friend-zva/golang-course-task2/collector/internal/usecase"
	pb "github.com/Friend-zva/golang-course-task2/collector/pkg/api/v1"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: (%s)", err)
	}

	clientGH := github.NewGitHubAPI(http.DefaultClient)

	info := usecase.NewInfoRepo(clientGH)

	serverGRPC := grpcS.NewServer(info)

	server := grpc.NewServer()
	pb.RegisterInfoRepoServiceServer(server, serverGRPC)

	log.Printf("Listening server at %s", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: (%s)", err)
	}
}
