package http

import "net/http"

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "Endpoint not found. Check /swagger/index.html for available routes"}`))
}
