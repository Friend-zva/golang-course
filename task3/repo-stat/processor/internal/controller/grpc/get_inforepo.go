package grpc

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/dto"
	processorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/processor"
)

func (h *Handler) GetInfoRepo(ctx context.Context, req *processorpb.GetInfoRepoRequest) (*processorpb.GetInfoRepoResponse, error) {
	input := dto.GetInfoRepoInput{
		Owner: req.Owner,
		Repo:  req.Repo,
	}

	output, err := h.getInfoRepo.Execute(ctx, input)
	if err != nil {
		var appErr *domain.AppError
		if errors.As(err, &appErr) {
			grpcCode := codes.Internal
			switch appErr.Code {
			case domain.CodeNotFound:
				grpcCode = codes.NotFound
			case domain.CodeInternal:
				grpcCode = codes.Internal
			case domain.CodeGateway:
				grpcCode = codes.Unavailable
			}
			return nil, status.Error(grpcCode, appErr.Message)
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &processorpb.GetInfoRepoResponse{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    timestamppb.New(time.Time(output.DateCreation)),
		CountStargazers: int32(output.CountStargazers),
		CountForks:      int32(output.CountForks),
	}, nil
}
