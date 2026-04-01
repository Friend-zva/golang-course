package http

import (
	"log/slog"
	"net/http"
)

func NewRouter(
	log *slog.Logger,
	subscriberPing ServicePing,
	processorPing ServicePing,
	getInfoRepo ProcessorGetInfoRepo,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /api/ping", NewPingHandler(log, subscriberPing, processorPing))
	mux.Handle("GET /api/repositories/info", NewInfoRepoHandler(log, getInfoRepo))
	mux.HandleFunc("/", NotFoundHandler(log))

	return mux
}
