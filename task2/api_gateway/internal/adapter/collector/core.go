package collector

import (
	"context"
	"fmt"
	"time"

	"github.com/Friend-zva/golang-course-task2/api_gateway/dto/driven"
	"github.com/Friend-zva/golang-course-task2/api_gateway/internal/domain"
	pb "github.com/Friend-zva/golang-course-task2/api_gateway/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CollectorAPI struct {
	client pb.InfoRepoServiceClient
}

func (c *CollectorAPI) GetInfoRepo(ctx context.Context, input driven.CollectorInput) (domain.InfoRepo, error) {
	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return domain.InfoRepo{}, fmt.Errorf("Failed to create connection (%w)", err)
	}
	defer conn.Close()

	client := pb.NewInfoRepoServiceClient(conn)

	req := pb.GetInfoRepoRequest{
		Owner: input.Owner,
		Repo:  input.Repo,
	}
	resp, err := client.GetInfoRepo(ctx, &req)
	if err != nil {
		return domain.InfoRepo{}, fmt.Errorf("Failed to create request (%w)", err)
	}

	dateCreation, err := time.Parse(time.RFC1123, resp.DateCreation)
	if err != nil {
		return domain.InfoRepo{}, fmt.Errorf("Failed to parse date (%w)", err)
	}

	return domain.InfoRepo{
		Name:            resp.Name,
		Description:     resp.Description,
		DateCreation:    dateCreation,
		CountStargazers: int(resp.CountStargazers),
		CountForks:      int(resp.CountForks),
	}, nil
}
