package http

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	dto "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/dto"
	pkg "github.com/Friend-zva/golang-course-task3/repo-stat/api/pkg"
	apperror "github.com/Friend-zva/golang-course-task3/repo-stat/platform/apperror"
)

type ErrorResponse struct {
	Error string `json:"error" example:"repo 'go-course' not found"`
}

type InfoRepoHandler struct {
	log     *slog.Logger
	usecase ProcessorGetInfoRepo
}

func NewInfoRepoHandler(log *slog.Logger, usecase ProcessorGetInfoRepo) *InfoRepoHandler {
	return &InfoRepoHandler{
		log:     log,
		usecase: usecase,
	}
}

// InfoRepoHandler gets information about a GitHub repository.
// @Summary Get repository info
// @Description Get information about a GitHub repository such as name, description, stars, forks, and creation date
// @Tags repository
// @Produce json
// @Param url query string true "Full URL of the GitHub repository" example("https://github.com/golang/go")
// @Success 200 {object} dto.GetInfoRepoOutput "Successful response with repository statistics"
// @Failure 400 {object} ErrorResponse "Invalid URL format"
// @Failure 404 {object} ErrorResponse "Repository not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Failure 502 {object} ErrorResponse "External GitHub API error"
// @Router /api/repositories/info [get]
func (iRH *InfoRepoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	owner, repo, err := pkg.ExtractPath(url)
	if err != nil {
		errResp := ErrorResponse{"invalid url format"}
		pkg.WriteJSON(*iRH.log, w, http.StatusBadRequest, errResp)
		return
	}

	output, err := iRH.usecase.Execute(r.Context(), owner, repo)
	if err != nil {
		var errApp *apperror.AppError
		if !errors.As(err, &errApp) {
			errApp = apperror.ErrInternal.Wrap(err)
		}

		errResp := ErrorResponse{errApp.Message}
		iRH.log.Error("cannot get inforepo", "error", errApp.Error())
		pkg.WriteJSON(*iRH.log, w, errApp.Code, errResp)
		return
	}

	resp := dto.GetInfoRepoResponse{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    output.DateCreation.Format(time.RFC3339),
		CountStargazers: output.CountStargazers,
		CountForks:      output.CountForks,
	}

	pkg.WriteJSON(*iRH.log, w, http.StatusOK, resp)
}
