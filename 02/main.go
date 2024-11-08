package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// main
func main() {
	fmt.Printf("change(\"LLRR\") = %s\n", change("LLRR"))
	fmt.Printf("change(\"==RLL\") = %s\n", change("==RLL"))
	fmt.Printf("change(\"=LLRR\") = %s\n", change("=LLRR"))
	fmt.Printf("change(\"RRL=R\") = %s\n", change("RRL=R"))
	fmt.Printf("change(\"LL-RR\") = %s\n", change("LL-RR"))
}

func change(encoded string) string {
	var (
		sequence = []int{}
		current  = 0
	)
	sequence = append(sequence, current)
	for i := 0; i < len(encoded); i++ {
		next := current
		if encoded[i] == 'L' {
			next = current - 1
		} else if encoded[i] == 'R' {
			next = current + 1
		} else if encoded[i] == '=' {
			next = current
		}
		sequence = append(sequence, next)
		current = next
	}
	log.Printf("sequence: %v", sequence)
	minValue := 0
	for _, num := range sequence {
		if num < minValue {
			minValue = num
		}
	}

	if minValue < 0 {
		offset := -minValue
		for i := range sequence {
			sequence[i] += offset
		}
	}
	var result strings.Builder
	for _, digit := range sequence {
		result.WriteString(strconv.Itoa(digit))
	}

	return result.String()
}
