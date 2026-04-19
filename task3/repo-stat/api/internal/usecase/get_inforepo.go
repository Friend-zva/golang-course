package usecase

import (
	"context"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
)

type GetInfoRepo struct {
	processor Processor
}

func NewGetInfoRepo(processor Processor) *GetInfoRepo {
	return &GetInfoRepo{
		processor: processor,
	}
}

func (gIR *GetInfoRepo) Execute(ctx context.Context, owner, repo string) (domain.InfoRepo, error) {
	info, err := gIR.processor.GetInfoRepo(ctx, owner, repo)
	if err != nil {
		return domain.InfoRepo{}, err
	}

	return info, nil
}
