package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Friend-zva/golang-course-task2/api_gateway/config"
	"github.com/Friend-zva/golang-course-task2/api_gateway/docs"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/adapter/collector"
	httpH "github.com/Friend-zva/golang-course-task2/api_gateway/internal/controller/http"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/usecase"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API Gateway
// @version 1.0.0
// @description API Server for getting GitHub repository information via gRPC Collector
// @BasePath /
func main() {
	cfg := config.MustLoad("config/local.yaml")

	docs.SwaggerInfo.Host = "localhost:8080" // Swagger inside browser wants to fetch from `localhost:8080`

	clientColl, err := collector.NewCollectorAPI(cfg.CollectorClient.Address)
	if err != nil {
		log.Fatalf("failed to init collector client: %s", err)
	}
	defer func() {
		if err := clientColl.Close(); err != nil {
			log.Printf("failed to close connection: %s\n", err)
		}
	}()

	info := usecase.NewInfoRepo(clientColl)

	handler := httpH.NewHandlers(info)

	router := chi.NewRouter()

	router.Get("/{owner}/{repo}", handler.GetInfoRepo)
	router.Get("/swagger/*", httpSwagger.WrapHandler)
	router.NotFound(handler.NotFound)

	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		log.Printf("Starting server on %s...", cfg.HTTPServer.Address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down API Gateway server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %s\n", err)
	}
	log.Println("API Gateway server exited properly")
}
