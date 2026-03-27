package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Friend-zva/golang-course-task3/repo-stat/api/config"
	"github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/adapter/subscriber"
	"github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/usecase"
)

func NewHandler(ctx context.Context, log *slog.Logger, cfg config.Config) (http.Handler, error) {
	subscriberClient, err := subscriber.NewClient(cfg.Services.Subscriber, log)
	if err != nil {
		log.Error("cannot init subscriber adapter", "error", err)
		return nil, err
	}

	pingUseCase := usecase.NewPing(subscriberClient)

	mux := http.NewServeMux()
	AddRoutes(mux, log, pingUseCase)

	var handler http.Handler = mux
	return handler, nil
}
