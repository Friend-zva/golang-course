package http

import (
	"errors"
	"net/http"

	"github.com/Friend-zva/golang-course-task2/api_gateway/dto/driving"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/domain"
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
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Failure 502 {object} ErrorResponse "Bad Gateway"
// @Failure 504 {object} ErrorResponse "Gateway Timeout"
// @Router /{owner}/{repo} [get]
func (h *Handler) GetInfoRepo(w http.ResponseWriter, r *http.Request) {
	input := driving.GetInfoRepoInput{
		Owner: chi.URLParam(r, "owner"),
		Repo:  chi.URLParam(r, "repo"),
	}

	output, err := h.infoRepo.GetInfoRepo(r.Context(), input)
	if err != nil {
		var appErr *domain.AppError
		if errors.As(err, &appErr) {
			render.Status(r, int(appErr.Code))
			render.JSON(w, r, ErrorResponse{Error: appErr.Message})
			return
		}

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrorResponse{Error: "internal server error"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, output)
}
