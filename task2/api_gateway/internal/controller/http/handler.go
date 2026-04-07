package http

import (
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/usecase"
)

type ErrorResponse struct {
	Error string `json:"error" example:"repo 'go-course' not found"`
}

type Handler struct {
	infoRepo *usecase.InfoRepo
}

func NewHandlers(infoRepo *usecase.InfoRepo) *Handler {
	return &Handler{
		infoRepo: infoRepo,
	}
}
