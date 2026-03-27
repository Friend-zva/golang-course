package usecase

import (
	"context"

	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/dto/driven"
	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/dto/driving"
)

func (iR *InfoRepo) GetInfoRepo(ctx context.Context, input driving.GetInfoRepoInput) (driving.GetInfoRepoOutput, error) {
	inputGH := driven.CollectorInput{Owner: input.Owner, Repo: input.Repo}

	info, err := iR.collector.GetInfoRepo(ctx, inputGH)
	if err != nil {
		return driving.GetInfoRepoOutput{}, err
	}

	return driving.GetInfoRepoOutput{
		Name:            info.Name,
		Description:     info.Description,
		DateCreation:    info.DateCreation,
		CountStargazers: info.CountStargazers,
		CountForks:      info.CountForks,
	}, nil
}
