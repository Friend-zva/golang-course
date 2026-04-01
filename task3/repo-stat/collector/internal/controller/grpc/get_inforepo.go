package grpc

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/dto"
	collectorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/collector"
)

type GitHubGetInfoRepo interface {
	Execute(ctx context.Context, input dto.GetInfoRepoInput) (dto.GetInfoRepoOutput, error)
}

type InfoRepoHandler struct {
	collectorpb.UnimplementedCollectorServer
	log     *slog.Logger
	usecase GitHubGetInfoRepo
}

func NewInfoRepoHandler(log *slog.Logger, usecase GitHubGetInfoRepo) *InfoRepoHandler {
	return &InfoRepoHandler{
		log:     log,
		usecase: usecase,
	}
}

func (iRH *InfoRepoHandler) GetInfoRepo(ctx context.Context, req *collectorpb.GetInfoRepoRequest) (*collectorpb.GetInfoRepoResponse, error) {
	input := dto.GetInfoRepoInput{
		Owner: req.Owner,
		Repo:  req.Repo,
	}

	output, err := iRH.usecase.Execute(ctx, input)
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

	return &collectorpb.GetInfoRepoResponse{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    timestamppb.New(output.DateCreation),
		CountStargazers: int32(output.CountStargazers),
		CountForks:      int32(output.CountForks),
	}, nil
}
