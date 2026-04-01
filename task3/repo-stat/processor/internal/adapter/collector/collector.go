package collector

import (
	"context"
	"log/slog"

	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"

	apperror "github.com/Friend-zva/golang-course-task3/repo-stat/platform/apperror"
	domain "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/processor/internal/dto"
	collectorpb "github.com/Friend-zva/golang-course-task3/repo-stat/proto/collector"
)

type Client struct {
	log  *slog.Logger
	conn *grpc.ClientConn
	pb   collectorpb.CollectorClient
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
		pb:   collectorpb.NewCollectorClient(conn),
	}, nil
}

func (c *Client) GetInfoRepo(ctx context.Context, input dto.CollectorGetInfoRepoInput) (domain.InfoRepo, error) {
	req := collectorpb.GetInfoRepoRequest{
		Owner: input.Owner,
		Repo:  input.Repo,
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
