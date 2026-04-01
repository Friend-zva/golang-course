package http

import (
	"context"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/dto"
)

type ServicePing interface {
	Execute(ctx context.Context) domain.PingStatus
}

type ProcessorGetInfoRepo interface {
	Execute(ctx context.Context, input dto.GetInfoRepoInput) (dto.GetInfoRepoOutput, error)
}
