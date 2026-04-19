package grpc

import (
	"context"
	"log/slog"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/domain"
	apperror "github.com/Friend-zva/golang-course-task3/repo-stat/platform/apperror"
	collectorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/collector"
)

type GitHubGetInfoRepo interface {
	Execute(ctx context.Context, owner, repo string) (domain.InfoRepo, error)
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
	output, err := iRH.usecase.Execute(ctx, req.Owner, req.Repo)
	if err != nil {
		return nil, apperror.Pack(err)
	}

	return &collectorpb.GetInfoRepoResponse{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    timestamppb.New(output.DateCreation),
		CountStargazers: int32(output.CountStargazers),
		CountForks:      int32(output.CountForks),
	}, nil
}
