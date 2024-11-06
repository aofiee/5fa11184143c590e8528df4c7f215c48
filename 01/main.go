package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type (
	Hard [][]int
)

const (
	uri = "https://raw.githubusercontent.com/7-solutions/backend-challenge/refs/heads/main/files/hard.json"
)

func main() {

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	var hard Hard
	err = json.Unmarshal(body, &hard)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	log.Printf("Max sum: %d", find(hard))
}

func find(h Hard) int {
	var max = 0
	for row := 1; row < len(h); row++ {
		for col := 0; col < len(h[row]); col++ {
			if col == 0 {
				h[row][col] += h[row-1][col]
			} else if col == len(h[row])-1 {
				h[row][col] += h[row-1][col-1]
			} else {
				h[row][col] += chkMax(h[row-1][col], h[row-1][col-1])
			}
		}
	}
	for _, value := range h[len(h)-1] {
		max = chkMax(max, value)
	}
	return max
}

func chkMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
