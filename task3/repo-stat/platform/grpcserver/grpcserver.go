package grpcserver

import (
	"context"
	"net"
	"time"

	grpc "google.golang.org/grpc"
)

type Config struct {
	Address    string `yaml:"address" env:"LISTEN_ADDRESS" env-default:"localhost:8081"`
	TimeoutSec int    `yaml:"timeout_sec" env:"TIMEOUT_SEC" env-default:"5"`
	Timeout    time.Duration
}

type Server struct {
	srv *grpc.Server
	lis net.Listener
}

func New(address string, opts ...grpc.ServerOption) (*Server, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Server{
		srv: grpc.NewServer(opts...),
		lis: lis,
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.srv.GracefulStop()
	}()
	return s.srv.Serve(s.lis)
}

func (s *Server) GRPC() *grpc.Server {
	return s.srv
}
