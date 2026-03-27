package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpcH "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/controller/grpc"
	processorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/processor"

	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/config"
	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/adapter/collector"
	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.MustLoad("config/config.yaml")

	listener, err := net.Listen("tcp", cfg.ProcessorServer.Address)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	clientColl, err := collector.NewCollectorAPI(cfg.CollectorClient.Address)
	if err != nil {
		log.Fatalf("failed to init collector client: %s", err)
	}
	defer func() {
		if err := clientColl.Close(); err != nil {
			log.Printf("failed to close connection: %s", err)
		}
	}()

	info := usecase.NewInfoRepo(clientColl)

	handler := grpcH.NewHandler(info)

	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(cfg.ProcessorServer.IdleTimeout),
	)
	processorpb.RegisterProcessorServer(grpcServer, handler)

	go func() {
		log.Printf("Starting server on %s...", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Printf("failed to serve: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Processor gRPC server...")

	grpcServer.GracefulStop()
	log.Println("Processor server exited properly")
}
