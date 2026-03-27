package usecase

import (
	"context"

	"github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
)

type SubscriberPinger interface {
	Ping(ctx context.Context) (domain.PingStatus, error)
}
