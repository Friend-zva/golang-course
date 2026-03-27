package domain

import (
	"time"
)

type InfoRepo struct {
	Name            string
	Description     string
	DateCreation    time.Time
	CountStargazers int
	CountForks      int
}
