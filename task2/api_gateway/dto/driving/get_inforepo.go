package driving

type GetInfoRepoInput struct {
	Owner string
	Repo  string
}

type GetInfoRepoOutput struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	DateCreation    string `json:"date_creation"`
	CountStargazers int    `json:"count_stargazers"`
	CountForks      int    `json:"count_forks"`
}
