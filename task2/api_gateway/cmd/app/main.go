package main

import (
	"log"
	"net/http"

	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/adapter/collector"
	httpH "github.com/Friend-zva/golang-course-task2/api_gateway/internal/controller/http"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	clientC := collector.CollectorAPI{}

	info := usecase.NewInfoRepo(&clientC)

	handler := httpH.NewHandlers(info)

	router := chi.NewRouter()
	router.Get("/{owner}/{repo}", handler.GetInfoRepo)

	log.Println("Listening server at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
