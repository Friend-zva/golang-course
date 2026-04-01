package usecase

import (
	"context"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/dto"
)

type GitHub interface {
	GetInfoRepo(ctx context.Context, input dto.GitHubGetInfoRepoInput) (domain.InfoRepo, error)
}

type GetInfoRepo struct {
	github GitHub
}

func NewGetInfoRepo(github GitHub) *GetInfoRepo {
	return &GetInfoRepo{
		github: github,
	}
}

func (gIR *GetInfoRepo) Execute(ctx context.Context, input dto.GetInfoRepoInput) (dto.GetInfoRepoOutput, error) {
	inputGH := dto.GitHubGetInfoRepoInput(input)

	info, err := gIR.github.GetInfoRepo(ctx, inputGH)
	if err != nil {
		return dto.GetInfoRepoOutput{}, err
	}

	return dto.GetInfoRepoOutput{
		InfoRepo: info,
	}, nil
}
