package grpc

import (
	collectorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/collector"

	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/usecase"
)

type Handler struct {
	collectorpb.UnimplementedCollectorServer
	infoRepo *usecase.InfoRepo
}

func NewHandler(iR *usecase.InfoRepo) *Handler {
	return &Handler{
		infoRepo: iR,
	}
}
