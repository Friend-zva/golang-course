package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/Friend-zva/golang-course-task2/api_gateway/dto/driven"
	"github.com/Friend-zva/golang-course-task2/api_gateway/dto/driving"
)

func (iR *InfoRepo) GetInfoRepo(ctx context.Context, input driving.GetInfoRepoInput) (driving.GetInfoRepoOutput, error) {
	inputGH := driven.CollectorInput{Owner: input.Owner, Repo: input.Repo}

	info, err := iR.collector.GetInfoRepo(ctx, inputGH)
	if err != nil {
		return driving.GetInfoRepoOutput{}, fmt.Errorf("Failed to get info: (%w)", err)
	}

	return driving.GetInfoRepoOutput{
		Name:            info.Name,
		Description:     info.Description,
		DateCreation:    info.DateCreation.Format(time.RFC1123),
		CountStargazers: info.CountStargazers,
		CountForks:      info.CountForks,
	}, nil
}
