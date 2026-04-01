package pkg

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func WriteJSON(log slog.Logger, w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error("cannot encode json", "error", err)
	}
}
