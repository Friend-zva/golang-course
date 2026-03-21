package usecase

import (
	"context"

	"github.com/Friend-zva/golang-course-task2/api_gateway/dto/driven"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/adapter/collector"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/domain"
)

type Collector interface {
	GetInfoRepo(ctx context.Context, input driven.CollectorInput) (domain.InfoRepo, error)
}

type InfoRepo struct {
	collector Collector
}

func NewInfoRepo(collector *collector.CollectorAPI) *InfoRepo {
	return &InfoRepo{
		collector: collector,
	}
}
