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
	client  pb.InfoRepoServiceClient
	address string
}

func NewCollectorAPI(address string) *CollectorAPI {
	return &CollectorAPI{
		address: address,
	}
}

func (c *CollectorAPI) GetInfoRepo(ctx context.Context, input driven.CollectorInput) (domain.InfoRepo, error) {
	conn, err := grpc.NewClient(c.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(fmt.Errorf("failed to create connection: %w", err))
	}
	defer conn.Close()

	client := pb.NewInfoRepoServiceClient(conn)

	req := pb.GetInfoRepoRequest{
		Owner: input.Owner,
		Repo:  input.Repo,
	}
	resp, err := client.GetInfoRepo(ctx, &req)
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
