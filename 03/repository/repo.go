package repository

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type (
	Repo struct {
		URI string
	}
)

func NewRepo(uri string) *Repo {
	return &Repo{
		URI: uri,
	}
}

func (r *Repo) GetSummary() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/?type=meat-and-filler&paras=99&format=text", r.URI))
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
		return "", err
	}

	text := string(body)
	return text, nil
}
