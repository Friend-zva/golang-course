package grpc

import (
	"context"

	processorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/processor"
)

func (h *Handler) Ping(ctx context.Context, _ *processorpb.PingRequest) (*processorpb.PingResponse, error) {
	h.log.Debug("processor ping request received")

	return &processorpb.PingResponse{
		Reply: h.ping.Execute(ctx),
	}, nil
}
