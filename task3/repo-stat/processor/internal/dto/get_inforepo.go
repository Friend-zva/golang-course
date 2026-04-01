package dto

import "time"

type GetInfoRepoInput struct {
	Owner string
	Repo  string
}

type GetInfoRepoOutput struct {
	Name            string
	Description     string
	DateCreation    time.Time
	CountStargazers int
	CountForks      int
}

type CollectorGetInfoRepoInput struct {
	Owner string
	Repo  string
}
