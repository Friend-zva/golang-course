package dto

type GetInfoRepoInput struct {
	Owner string
	Repo  string
}

type GetInfoRepoOutput struct {
	Name            string `json:"full_name" example:"golang-course"`
	Description     string `json:"description" example:"Homework for GoLang course 2026"`
	DateCreation    string `json:"created_at" example:"Fri, 20 Aug 2021 09:38:00 UTC"`
	CountStargazers int    `json:"stars" example:"52"`
	CountForks      int    `json:"forks" example:"45"`
}

type ProcessorGetRepoInfoInput struct {
	Owner string
	Repo  string
}
