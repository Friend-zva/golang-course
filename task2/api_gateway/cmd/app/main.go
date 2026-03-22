package main

import (
	"log"
	"net/http"

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

	docs.SwaggerInfo.Host = cfg.HTTPServer.Address

	clientColl := collector.NewCollectorAPI(cfg.CollectorClient.Address)

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

	log.Printf("starting server on %s", cfg.HTTPServer.Address)
	log.Fatal(server.ListenAndServe())
}
