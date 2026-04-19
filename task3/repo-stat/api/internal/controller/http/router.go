package http

import (
	"context"
	"log/slog"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Friend-zva/golang-course-task3/repo-stat/api/docs"
	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
)

type ServicePing interface {
	Execute(ctx context.Context) domain.PingStatus
}

type ProcessorGetInfoRepo interface {
	Execute(ctx context.Context, owner, repo string) (domain.InfoRepo, error)
}

func NewRouter(
	log *slog.Logger,
	subscriberPing ServicePing,
	processorPing ServicePing,
	getInfoRepo ProcessorGetInfoRepo,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /api/ping", NewPingHandler(log, subscriberPing, processorPing))
	mux.Handle("GET /api/repositories/info", NewInfoRepoHandler(log, getInfoRepo))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("/", NotFoundHandler(log))

	return mux
}
