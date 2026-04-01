package http

import (
	"log/slog"
	"net/http"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/dto"
	pkg "github.com/Friend-zva/golang-course-task3/repo-stat/api/pkg"
)

type PingHandler struct {
	log            *slog.Logger
	subscriberPing ServicePing
	processorPing  ServicePing
}

func NewPingHandler(log *slog.Logger, sub ServicePing, proc ServicePing) *PingHandler {
	return &PingHandler{
		log:            log,
		subscriberPing: sub,
		processorPing:  proc,
	}
}

func (pH *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	subStatus := pH.subscriberPing.Execute(r.Context())
	procStatus := pH.processorPing.Execute(r.Context())

	response := dto.PingResponse{
		Status: "ok",
		Services: []dto.ServiceStatus{
			{
				Name:   "processor",
				Status: string(procStatus),
			},
			{
				Name:   "subscriber",
				Status: string(subStatus),
			},
		},
	}

	statusCode := http.StatusOK
	if subStatus != domain.PingStatusUp || procStatus != domain.PingStatusUp {
		response.Status = "degraded"
		statusCode = http.StatusServiceUnavailable
	}

	pkg.WriteJSON(*pH.log, w, statusCode, response)
}
