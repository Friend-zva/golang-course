package http

import (
	"net/http"

	"github.com/Friend-zva/golang-course-task2/api_gateway/dto/driving"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handlers) GetInfoRepo(w http.ResponseWriter, r *http.Request) {
	input := driving.GetInfoRepoInput{
		Owner: chi.URLParam(r, "owner"),
		Repo:  chi.URLParam(r, "repo"),
	}

	output, err := h.infoRepo.GetInfoRepo(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, output)
}
