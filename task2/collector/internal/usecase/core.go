package usecase

import (
	"context"

	"github.com/Friend-zva/golang-course-task2/collector/dto/driven"
	"github.com/Friend-zva/golang-course-task2/collector/internal/adapter/github"
	"github.com/Friend-zva/golang-course-task2/collector/internal/domain"
)

type GitHub interface {
	GetInfoRepo(ctx context.Context, input driven.GitHubRepoInput) (domain.InfoRepo, error)
}

type InfoRepo struct {
	github GitHub
}

func NewInfoRepo(github *github.GitHubAPI) *InfoRepo {
	return &InfoRepo{
		github: github,
	}
}
