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
	Message         string    `json:"message"`
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	info := InfoRepo{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Fatal(err)
	}

	if info.Message != "" {
		fmt.Printf("client failed: %s\n", info.Message)
		os.Exit(1)
	}

	fmt.Println("Repository    :", repo)
	fmt.Println("Description   :", info.Description)
	fmt.Println("Created  at   :", info.DateCreation.Format(time.RFC1123))
	fmt.Println("Count of stars:", info.CountStargazers)
	fmt.Println("Count of forks:", info.CountForks)
}
