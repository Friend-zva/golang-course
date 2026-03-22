package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, ErrorResponse{
		Error: "Endpoint not found. Check /swagger/index.html for available routes",
	})
}
