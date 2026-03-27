package grpc

import (
	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/usecase"
	processorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/processor"
)

type Handler struct {
	processorpb.UnimplementedProcessorServer
	infoRepo *usecase.InfoRepo
}

func NewHandler(iR *usecase.InfoRepo) *Handler {
	return &Handler{
		infoRepo: iR,
	}
}
