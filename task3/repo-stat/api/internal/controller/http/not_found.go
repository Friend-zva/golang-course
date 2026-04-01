package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func NotFoundHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		errResp := ErrorResponse{
			Error: "Endpoint not found. Check /swagger/index.html for available routes",
		}

		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			log.Error("cannot encode not found response", "error", err)
		}
	}
}
