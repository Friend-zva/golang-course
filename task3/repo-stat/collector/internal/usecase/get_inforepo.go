package usecase

import (
	"context"

	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/dto/driven"
	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/dto/driving"
)

func (iR *InfoRepo) GetInfoRepo(ctx context.Context, input driving.GetInfoRepoInput) (driving.GetInfoRepoOutput, error) {
	inputGH := driven.GitHubRepoInput{Owner: input.Owner, Repo: input.Repo}

	info, err := iR.github.GetInfoRepo(ctx, inputGH)
	if err != nil {
		return driving.GetInfoRepoOutput{}, err
	}

	return driving.GetInfoRepoOutput{
		InfoRepo: info,
	}, nil
}
