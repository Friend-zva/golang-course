package http

import (
	"net/http"

	"github.com/Friend-zva/golang-course-task2/api_gateway/dto/driving"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// GetInfoRepo gets information about a GitHub repository.
// @Summary Get repository info
// @Description Get information about a GitHub repository such as name, description, stars, forks, and creation date
// @Tags repository
// @Produce json
// @Param owner path string true "GitHub username or organization name" example("suvorovrain")
// @Param repo path string true "Repository name" example("golang-course")
// @Success 200 {object} driving.GetInfoRepoOutput "Successful response with repository info"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /{owner}/{repo} [get]
func (h *Handler) GetInfoRepo(w http.ResponseWriter, r *http.Request) {
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
