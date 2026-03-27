package usecase

import (
	"context"

	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/adapter/collector"
	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/domain"
	"github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/dto/driven"
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
