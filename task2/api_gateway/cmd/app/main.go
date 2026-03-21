package main

import (
	"log"
	"net/http"

	"github.com/Friend-zva/golang-course-task2/api_gateway/config"
	_ "github.com/Friend-zva/golang-course-task2/api_gateway/docs"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/adapter/collector"
	httpH "github.com/Friend-zva/golang-course-task2/api_gateway/internal/controller/http"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/usecase"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API Gateway
// @version 1.0.0
// @description API Server for getting GitHub repository information via gRPC Collector
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.MustLoad("config/local.yaml")

	clientColl := collector.CollectorAPI{}

	info := usecase.NewInfoRepo(&clientColl)

	handler := httpH.NewHandlers(info)

	router := chi.NewRouter()
	router.Get("/{owner}/{repo}", handler.GetInfoRepo)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	log.Printf("Starting server on %s", cfg.HTTPServer.Address)
	log.Fatal(server.ListenAndServe())

}
