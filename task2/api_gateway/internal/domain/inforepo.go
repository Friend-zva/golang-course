package domain

import (
	"fmt"
	"time"
)

type InfoRepo struct {
	Name            string
	Description     string
	DateCreation    time.Time
	CountStargazers int
	CountForks      int
}

func (iR InfoRepo) String() string {
	return fmt.Sprintf(`=== %s ===
Description    : %s
Creation date  : %s
Count of stars : %d
Count of forks : %d`,
		iR.Name,
		iR.Description,
		iR.DateCreation.Format(time.RFC1123),
		iR.CountStargazers,
		iR.CountForks,
	)
}
