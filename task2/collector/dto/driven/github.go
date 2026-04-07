package driven

import "time"

type GitHubRepoInput struct {
	Owner string
	Repo  string
}

type GitHubRepoOutput struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	DateCreation    time.Time `json:"created_at"`
	CountStargazers int       `json:"stargazers_count"`
	CountForks      int       `json:"forks"`
}
