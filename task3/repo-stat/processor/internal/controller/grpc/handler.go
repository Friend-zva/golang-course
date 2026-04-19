package grpc

import (
	"context"
	"log/slog"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/domain"
	processorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/processor"
)

type Ping interface {
	Execute(ctx context.Context) string
}

type CollectorGetInfoRepo interface {
	Execute(ctx context.Context, owner, repo string) (domain.InfoRepo, error)
}

type Handler struct {
	processorpb.UnimplementedProcessorServer
	log         *slog.Logger
	ping        Ping
	getInfoRepo CollectorGetInfoRepo
}

func NewHandler(log *slog.Logger, ping Ping, getInfoRepo CollectorGetInfoRepo) *Handler {
	return &Handler{
		log:         log,
		ping:        ping,
		getInfoRepo: getInfoRepo,
	}
}
