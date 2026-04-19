package usecase

import (
	"context"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
)

type Pinger interface {
	Ping(ctx context.Context) domain.PingStatus
}

type Processor interface {
	GetInfoRepo(ctx context.Context, owner, repo string) (domain.InfoRepo, error)
}
