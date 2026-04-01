package http

import (
	"encoding/json"
	"log/slog"
	"net/http"

	dto "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/dto"
	pkg "github.com/Friend-zva/golang-course-task3/repo-stat/api/pkg"
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
		http.Error(w, "invalid url format", http.StatusBadRequest)
		return
	}

	input := dto.GetInfoRepoInput{
		Owner: owner,
		Repo:  repo,
	}

	output, err := iRH.usecase.Execute(r.Context(), input)
	if err != nil {
		iRH.log.Error("cannot get info repo", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(output); err != nil {
		iRH.log.Error("cannot encode inforepo response", "error", err)
	}
}
