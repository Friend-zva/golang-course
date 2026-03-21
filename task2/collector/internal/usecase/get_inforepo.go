package usecase

import (
	"context"
	"fmt"

	"github.com/Friend-zva/golang-course-task2/collector/dto/driven"
	"github.com/Friend-zva/golang-course-task2/collector/dto/driving"
)

func (iR *InfoRepo) GetInfoRepo(ctx context.Context, input driving.GetInfoRepoInput) (driving.GetInfoRepoOutput, error) {
	inputGH := driven.GitHubRepoInput{Owner: input.Owner, Repo: input.Repo}

	info, err := iR.github.GetInfoRepo(ctx, inputGH)
	if err != nil {
		return driving.GetInfoRepoOutput{}, fmt.Errorf("Failed to get info: (%w)", err)
	}

	return driving.GetInfoRepoOutput{
		InfoRepo: info,
	}, nil
}
