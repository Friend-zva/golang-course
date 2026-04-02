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

// PingHandler gets information about status Processor & Subscriber service.
// @Summary Health Check & Service Discovery
// @Description Monitors the availability of the entire internal chain (Processor and Subscriber).
// @Description Returns 200 OK if all services are reachable, or 503 Service Unavailable if any service is down.
// @Tags status
// @Produce json
// @Success 200 {object} dto.PingResponse "Successful response (All services UP)"
// @Failure 503 {object} dto.PingResponse "Service Unavailable (System Degraded)"
// @Router /api/ping [get]
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
