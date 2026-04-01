package http

import (
	"errors"
	"log/slog"
	"net/http"

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

func (iRH *InfoRepoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	owner, repo, err := pkg.ExtractPath(url)
	if err != nil {
		errResp := ErrorResponse{"invalid url format"}
		pkg.WriteJSON(*iRH.log, w, http.StatusBadRequest, errResp)
		return
	}

	input := dto.GetInfoRepoInput{
		Owner: owner,
		Repo:  repo,
	}

	output, err := iRH.usecase.Execute(r.Context(), input)
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

	pkg.WriteJSON(*iRH.log, w, http.StatusOK, output)
}
