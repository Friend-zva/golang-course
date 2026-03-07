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
	CountStargazers int       `json:"stargazers_count"`
	CountForks      int       `json:"forks"`
	DateCreation    time.Time `json:"created_at"`
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("args: incorrect repository entry")
	}

	repo := args[1]
	url := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		log.Fatalf("client failed: repo '%s' not found\n", repo)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("client failed: %s\n", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	info := InfoRepo{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Repository     :", repo)
	fmt.Println("Description    :", info.Description)
	fmt.Println("Created  at    :", info.DateCreation.Format(time.RFC1123))
	fmt.Println("Count of stars :", info.CountStargazers)
	fmt.Println("Count of forks :", info.CountForks)
}
