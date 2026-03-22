package controller

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Friend-zva/golang-course-task2/collector/dto/driving"
	"github.com/Friend-zva/golang-course-task2/collector/internal/domain"
	pb "github.com/Friend-zva/golang-course-task2/proto/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetInfoRepo(ctx context.Context, req *pb.GetInfoRepoRequest) (*pb.GetInfoRepoResponse, error) {
	input := driving.GetInfoRepoInput{
		Owner: req.Owner,
		Repo:  req.Repo,
	}

	output, err := h.infoRepo.GetInfoRepo(ctx, input)
	if err != nil {
		var appErr *domain.AppError
		if errors.As(err, &appErr) {
			grpcCode := codes.Internal
			switch appErr.Code {
			case domain.RepoNotFound:
				grpcCode = codes.NotFound
			case domain.CodeInternal:
				grpcCode = codes.Internal
			case domain.CodeExternal:
				grpcCode = codes.Unavailable
			}
			return nil, status.Error(grpcCode, appErr.Message)
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetInfoRepoResponse{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    timestamppb.New(output.DateCreation),
		CountStargazers: int32(output.CountStargazers),
		CountForks:      int32(output.CountForks),
	}, nil
}
