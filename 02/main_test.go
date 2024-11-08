package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestChange(t *testing.T) {
	t.Run("Change", func(t *testing.T) {
		s := change("LLRR")
		assert.Equal(t, s, "210122")
	})

	t.Run("Change", func(t *testing.T) {
		s := change("==RLL")
		assert.Equal(t, s, "000210")
	})

	t.Run("Change", func(t *testing.T) {
		s := change("=LLRR")
		assert.Equal(t, s, "221012")
	})

	t.Run("Change", func(t *testing.T) {
		s := change("RRL=R")
		assert.Equal(t, s, "012001")
	})
}
