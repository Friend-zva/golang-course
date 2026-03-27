package usecase

import (
	"context"

	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/adapter/github"
	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/domain"
	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/dto/driven"
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
