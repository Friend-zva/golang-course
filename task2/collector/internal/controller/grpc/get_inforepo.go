package controller

import (
	"context"
	"time"

	"github.com/Friend-zva/golang-course-task2/collector/dto/driving"
	pb "github.com/Friend-zva/golang-course-task2/proto/pkg/api/v1"
)

func (h *Handler) GetInfoRepo(ctx context.Context, req *pb.GetInfoRepoRequest) (*pb.GetInfoRepoResponse, error) {
	input := driving.GetInfoRepoInput{
		Owner: req.Owner,
		Repo:  req.Repo,
	}

	output, err := h.infoRepo.GetInfoRepo(ctx, input)
	if err != nil {
		return &pb.GetInfoRepoResponse{}, err
	}

	return &pb.GetInfoRepoResponse{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    output.DateCreation.Format(time.RFC1123),
		CountStargazers: int32(output.CountStargazers),
		CountForks:      int32(output.CountForks),
	}, nil
}
