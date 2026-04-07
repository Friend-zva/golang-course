package controller

import (
	"github.com/Friend-zva/golang-course-task2/collector/internal/usecase"
	pb "github.com/Friend-zva/golang-course-task2/proto/pkg/api/v1"
)

type Handler struct {
	pb.UnimplementedInfoRepoServiceServer
	infoRepo *usecase.InfoRepo
}

func NewHandler(iR *usecase.InfoRepo) *Handler {
	return &Handler{
		infoRepo: iR,
	}
}
