package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type InfoRepo struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	DateCreation    time.Time `json:"created_at"`
	CountStargazers int       `json:"stargazers_count"`
	CountForks      int       `json:"forks"`
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

func extractPath(link string) (string, error) {
	endpoints := strings.Split(link, "/")
	repo := ""
	var err error

	switch len(endpoints) {
	case 2: // owner/repo
		repo = fmt.Sprintf("%s/%s", endpoints[0], endpoints[1])
	case 5: // https://github.com/owner/repo
		repo = fmt.Sprintf("%s/%s", endpoints[3], endpoints[4])
	default:
		err = fmt.Errorf("extract path: incorrect link")
	}

	return repo, err
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal(`
Usage:
> golang-course-task1 <owner/repo>
> golang-course-task1 https://github.com/<owner/repo>`)
	}

	repo, err := extractPath(args[1])
	if err != nil {
		log.Fatalf("Failed to validate arg (%s)\n", err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get response (%s)\n", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close body response (%s)\n", err)
		}
	}()

	switch resp.StatusCode {
	case http.StatusOK:
		break
	case http.StatusNotFound:
		log.Fatalf("Repo '%s' not found\n", repo)
	default:
		log.Fatalf("Client failed: %s\n", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body response (%s)\n", err)
	}

	info := InfoRepo{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Fatalf("Failed to serialize data (%s)\n", err)
	}

	fmt.Println(info)
}
