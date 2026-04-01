package usecase

import (
	"context"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/dto"
)

type Collector interface {
	GetInfoRepo(ctx context.Context, input dto.CollectorGetInfoRepoInput) (domain.InfoRepo, error)
}

type GetInfoRepo struct {
	collector Collector
}

func NewGetInfoRepo(collector Collector) *GetInfoRepo {
	return &GetInfoRepo{
		collector: collector,
	}
}

func (gIR *GetInfoRepo) Execute(ctx context.Context, input dto.GetInfoRepoInput) (dto.GetInfoRepoOutput, error) {
	inputGH := dto.CollectorGetInfoRepoInput(input)

	info, err := gIR.collector.GetInfoRepo(ctx, inputGH)
	if err != nil {
		return dto.GetInfoRepoOutput{}, err
	}

	return dto.GetInfoRepoOutput(info), nil
}
