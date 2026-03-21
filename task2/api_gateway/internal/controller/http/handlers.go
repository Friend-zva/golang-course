package http

import (
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/usecase"
)

type Handlers struct {
	infoRepo *usecase.InfoRepo
}

func NewHandlers(infoRepo *usecase.InfoRepo) *Handlers {
	return &Handlers{
		infoRepo: infoRepo,
	}
}
