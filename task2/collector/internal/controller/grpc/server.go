package controller

import (
	"github.com/Friend-zva/golang-course-task2/collector/internal/usecase"
	pb "github.com/Friend-zva/golang-course-task2/proto/pkg/api/v1"
)

type Server struct {
	pb.UnimplementedInfoRepoServiceServer
	infoRepo *usecase.InfoRepo
}

func NewServer(iR *usecase.InfoRepo) *Server {
	return &Server{
		infoRepo: iR,
	}
}
