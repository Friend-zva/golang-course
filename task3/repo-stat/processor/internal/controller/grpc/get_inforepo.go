package grpc

import (
	"context"
	"time"

	apperror "github.com/Friend-zva/golang-course-task3/repo-stat/platform/apperror"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/dto"
	processorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/processor"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetInfoRepo(ctx context.Context, req *processorpb.GetInfoRepoRequest) (*processorpb.GetInfoRepoResponse, error) {
	input := dto.GetInfoRepoInput{
		Owner: req.Owner,
		Repo:  req.Repo,
	}

	output, err := h.getInfoRepo.Execute(ctx, input)
	if err != nil {
		return nil, apperror.Pack(err)
	}

	return &processorpb.GetInfoRepoResponse{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    timestamppb.New(time.Time(output.DateCreation)),
		CountStargazers: int32(output.CountStargazers),
		CountForks:      int32(output.CountForks),
	}, nil
}
