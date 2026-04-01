package processor

import (
	"context"
	"log/slog"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	insecure "google.golang.org/grpc/credentials/insecure"
	status "google.golang.org/grpc/status"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/api/internal/dto"
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

func (c *Client) GetInfoRepo(ctx context.Context, input dto.ProcessorGetRepoInfoInput) (domain.InfoRepo, error) {
	req := processorpb.GetInfoRepoRequest{
		Owner: input.Owner,
		Repo:  input.Repo,
	}

	resp, err := c.pb.GetInfoRepo(ctx, &req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return domain.InfoRepo{}, domain.ErrNotFound.WithMessage(st.Message())
			case codes.Unavailable:
				return domain.InfoRepo{}, domain.ErrGateway.WithMessage(st.Message())
			case codes.DeadlineExceeded:
				return domain.InfoRepo{}, domain.ErrTimeout.WithMessage(st.Message())
			}
		}
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(err)
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
