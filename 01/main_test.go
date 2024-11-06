package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		input := `[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]`
		var h Hard
		err := json.Unmarshal([]byte(input), &h)
		if err != nil {
			log.Fatalf("Failed to unmarshal response body: %v", err)
		}
		got := find(h)
		want := 237
		assert.Equal(t, want, got)
	})

	t.Run("Test case 2", func(t *testing.T) {
		input := `[[1], [2, 3], [4, 5, 6]]`
		var h Hard
		err := json.Unmarshal([]byte(input), &h)
		if err != nil {
			log.Fatalf("Failed to unmarshal response body: %v", err)
		}
		got := find(h)
		want := 10
		assert.Equal(t, want, got)
	})
}
