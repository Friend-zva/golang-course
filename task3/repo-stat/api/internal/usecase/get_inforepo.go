package usecase

import (
	"context"
	"time"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/dto"
)

type Processor interface {
	GetInfoRepo(ctx context.Context, input dto.ProcessorGetRepoInfoInput) (domain.InfoRepo, error)
}

type GetInfoRepo struct {
	processor Processor
}

func NewGetInfoRepo(processor Processor) *GetInfoRepo {
	return &GetInfoRepo{
		processor: processor,
	}
}

func (gIR *GetInfoRepo) Execute(ctx context.Context, input dto.GetInfoRepoInput) (dto.GetInfoRepoOutput, error) {
	inputProc := dto.ProcessorGetRepoInfoInput(input)

	info, err := gIR.processor.GetInfoRepo(ctx, inputProc)
	if err != nil {
		return dto.GetInfoRepoOutput{}, err
	}

	return dto.GetInfoRepoOutput{
		Name:            info.Name,
		Description:     info.Description,
		DateCreation:    info.DateCreation.Format(time.RFC1123),
		CountStargazers: info.CountStargazers,
		CountForks:      info.CountForks,
	}, nil
}
