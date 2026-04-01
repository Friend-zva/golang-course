package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	domain "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/domain"
	dto "github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/dto"
)

type Client struct {
	log    *slog.Logger
	client *http.Client
}

func NewClient(client *http.Client, log *slog.Logger) *Client {
	return &Client{
		log:    log,
		client: client,
	}
}

type githubGetInfoRepoOutput struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	DateCreation    time.Time `json:"created_at"`
	CountStargazers int       `json:"stargazers_count"`
	CountForks      int       `json:"forks"`
}

func (c *Client) GetInfoRepo(ctx context.Context, input dto.GitHubGetInfoRepoInput) (domain.InfoRepo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", input.Owner, input.Repo)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(fmt.Errorf("cannot create request: %w", err))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return domain.InfoRepo{}, domain.ErrExternal.Wrap(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.log.Warn("cannot close body response", "error", err)
		}
	}()

	switch resp.StatusCode {
	case http.StatusOK:
		break
	case http.StatusNotFound:
		return domain.InfoRepo{}, domain.ErrNotFound.WithMessage(fmt.Sprintf("repo '%s' not found", input.Repo))
	default:
		return domain.InfoRepo{}, domain.ErrExternal.WithMessage(fmt.Sprintf("github api returned status: %s", resp.Status))
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(fmt.Errorf("cannot read body response: %w", err))
	}

	output := githubGetInfoRepoOutput{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(fmt.Errorf("cannot serialize data: %w", err))
	}

	return domain.InfoRepo{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    output.DateCreation,
		CountStargazers: output.CountStargazers,
		CountForks:      output.CountForks,
	}, nil
}
