package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/domain"
	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/dto/driven"
)

type GitHubAPI struct {
	client *http.Client
}

func NewGitHubAPI(client *http.Client) *GitHubAPI {
	if client == nil {
		client = http.DefaultClient
	}
	return &GitHubAPI{
		client: client,
	}
}

func (g *GitHubAPI) GetInfoRepo(ctx context.Context, input driven.GitHubRepoInput) (domain.InfoRepo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", input.Owner, input.Repo)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(fmt.Errorf("failed to create request: %w", err))
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return domain.InfoRepo{}, domain.ErrExternal.Wrap(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close body response: %s", err)
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
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(fmt.Errorf("failed to read body response: %w", err))
	}

	output := driven.GitHubRepoOutput{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return domain.InfoRepo{}, domain.ErrInternal.Wrap(fmt.Errorf("failed to serialize data: %w", err))
	}

	return domain.InfoRepo{
		Name:            output.Name,
		Description:     output.Description,
		DateCreation:    output.DateCreation,
		CountStargazers: output.CountStargazers,
		CountForks:      output.CountForks,
	}, nil
}
