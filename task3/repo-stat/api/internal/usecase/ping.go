package usecase

import (
	"context"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
)

type Ping struct {
	pinger Pinger
}

func NewPing(pinger Pinger) *Ping {
	return &Ping{
		pinger: pinger,
	}
}

func (p *Ping) Execute(ctx context.Context) domain.PingStatus {
	return p.pinger.Ping(ctx)
}
