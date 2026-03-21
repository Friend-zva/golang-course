package driving

type GetInfoRepoInput struct {
	Owner string
	Repo  string
}

// GetInfoRepoOutput represents the repository data returned to the user
type GetInfoRepoOutput struct {
	Name            string `json:"name" example:"golang-course"`
	Description     string `json:"description" example:"Homework for GoLang course 2026"`
	DateCreation    string `json:"date_creation" example:"Fri, 20 Aug 2021 09:38:00 UTC"`
	CountStargazers int    `json:"count_stargazers" example:"52"`
	CountForks      int    `json:"count_forks" example:"45"`
}
