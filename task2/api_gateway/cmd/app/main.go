package main

import (
	"log"
	"net/http"

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
	clientC := collector.CollectorAPI{}

	info := usecase.NewInfoRepo(&clientC)

	handler := httpH.NewHandlers(info)

	router := chi.NewRouter()
	router.Get("/{owner}/{repo}", handler.GetInfoRepo)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println("Listening server at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
