package http

import (
	"log/slog"
	"net/http"

	pkg "github.com/Friend-zva/golang-course-task3/repo-stat/api/pkg"
)

func NotFoundHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errResp := ErrorResponse{
			"Endpoint not found. Check /swagger/index.html for available routes",
		}
		pkg.WriteJSON(*log, w, http.StatusNotFound, errResp)
	}
}
