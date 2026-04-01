package http

import (
	"context"
	"log/slog"
	"net/http"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/dto"
)

type ServicePing interface {
	Execute(ctx context.Context) domain.PingStatus
}

type ProcessorGetInfoRepo interface {
	Execute(ctx context.Context, input dto.GetInfoRepoInput) (dto.GetInfoRepoOutput, error)
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
	mux.HandleFunc("/", NotFoundHandler(log))

	return mux
}
