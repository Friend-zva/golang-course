package driving

import (
	"github.com/Friend-zva/golang-course-task3/repo-stat/collector/internal/domain"
)

type GetInfoRepoInput struct {
	Owner string
	Repo  string
}

type GetInfoRepoOutput struct {
	domain.InfoRepo
}
