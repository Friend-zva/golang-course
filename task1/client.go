package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type InfoRepo struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	DateCreation    time.Time `json:"created_at"`
	CountStargazers int       `json:"stargazers_count"`
	CountForks      int       `json:"forks"`
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: golang-course-task1 <owner/repo>")
	}

	repo := args[1]
	url := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close body response: %s\n", err)
		}
	}()

	if resp.StatusCode == http.StatusNotFound {
		log.Fatalf("Repo '%s' not found\n", repo)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Client failed: %s\n", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body response: %s\n", err)
	}

	info := InfoRepo{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Fatalf("Failed to serialize data: %s\n", err)
	}

	fmt.Println("Repository     :", repo)
	fmt.Println("Description    :", info.Description)
	fmt.Println("Creation date  :", info.DateCreation.Format(time.RFC1123))
	fmt.Println("Count of stars :", info.CountStargazers)
	fmt.Println("Count of forks :", info.CountForks)
}
