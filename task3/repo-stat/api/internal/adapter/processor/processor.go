package processor

import (
	"context"
	"log/slog"

	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
	apperror "github.com/Friend-zva/golang-course-task3/repo-stat/platform/apperror"
	processorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/processor"
)

type Client struct {
	log  *slog.Logger
	conn *grpc.ClientConn
	pb   processorpb.ProcessorClient
}

func NewClient(address string, log *slog.Logger) (*Client, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		log:  log,
		conn: conn,
		pb:   processorpb.NewProcessorClient(conn),
	}, nil
}

func (c *Client) Ping(ctx context.Context) domain.PingStatus {
	_, err := c.pb.Ping(ctx, &processorpb.PingRequest{})
	if err != nil {
		c.log.Error("cannot ping processor", "error", err)
		return domain.PingStatusDown
	}

	return domain.PingStatusUp
}

func (c *Client) GetInfoRepo(ctx context.Context, owner, repo string) (domain.InfoRepo, error) {
	req := processorpb.GetInfoRepoRequest{
		Owner: owner,
		Repo:  repo,
	}

	resp, err := c.pb.GetInfoRepo(ctx, &req)
	if err != nil {
		return domain.InfoRepo{}, apperror.Unpack(err)
	}

	return domain.InfoRepo{
		Name:            resp.Name,
		Description:     resp.Description,
		DateCreation:    resp.DateCreation.AsTime(),
		CountStargazers: int(resp.CountStargazers),
		CountForks:      int(resp.CountForks),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
