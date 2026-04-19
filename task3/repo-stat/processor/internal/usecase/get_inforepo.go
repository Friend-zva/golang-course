package usecase

import (
	"context"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/domain"
)

type Collector interface {
	GetInfoRepo(ctx context.Context, owner, repo string) (domain.InfoRepo, error)
}

type GetInfoRepo struct {
	collector Collector
}

func NewGetInfoRepo(collector Collector) *GetInfoRepo {
	return &GetInfoRepo{
		collector: collector,
	}
}

func (gIR *GetInfoRepo) Execute(ctx context.Context, owner, repo string) (domain.InfoRepo, error) {
	info, err := gIR.collector.GetInfoRepo(ctx, owner, repo)
	if err != nil {
		return domain.InfoRepo{}, err
	}

	return info, nil
}
