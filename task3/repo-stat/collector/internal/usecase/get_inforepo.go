package usecase

import (
	"context"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/domain"
)

type GitHub interface {
	GetInfoRepo(ctx context.Context, owner, repo string) (domain.InfoRepo, error)
}

type GetInfoRepo struct {
	github GitHub
}

func NewGetInfoRepo(github GitHub) *GetInfoRepo {
	return &GetInfoRepo{
		github: github,
	}
}

func (gIR *GetInfoRepo) Execute(ctx context.Context, owner, repo string) (domain.InfoRepo, error) {
	info, err := gIR.github.GetInfoRepo(ctx, owner, repo)
	if err != nil {
		return domain.InfoRepo{}, err
	}

	return info, nil
}
