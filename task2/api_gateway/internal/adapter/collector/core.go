package collector

import (
	"context"
	"fmt"
	"time"

	"github.com/Friend-zva/golang-course-task2/api_gateway/dto/driven"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/domain"
	pb "github.com/Friend-zva/golang-course-task2/proto/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type CollectorAPI struct {
	client pb.InfoRepoServiceClient
	conn   *grpc.ClientConn
}

func NewCollectorAPI(address string) (*CollectorAPI, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	client := pb.NewInfoRepoServiceClient(conn)

	return &CollectorAPI{
		client: client,
		conn:   conn,
	}, nil
}

func (c *CollectorAPI) Close() error {
	return c.conn.Close()
}

func (c *CollectorAPI) GetInfoRepo(ctx context.Context, input driven.CollectorInput) (domain.InfoRepo, error) {
	req := pb.GetInfoRepoRequest{
		Owner: input.Owner,
		Repo:  input.Repo,
	}

	resp, err := c.client.GetInfoRepo(ctx, &req)
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

	dateCreation, err := time.Parse(time.RFC1123, resp.DateCreation)
	if err != nil {
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(fmt.Errorf("failed to parse date: %w", err))
	}

	return domain.InfoRepo{
		Name:            resp.Name,
		Description:     resp.Description,
		DateCreation:    dateCreation,
		CountStargazers: int(resp.CountStargazers),
		CountForks:      int(resp.CountForks),
	}, nil
}
