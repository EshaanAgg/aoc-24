package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("../data/day04/sample.txt")
	input := string(b)

	assert.Equal(t, 18, Part1(input))
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("../data/day04/sample.txt")
	input := string(b)

	assert.Equal(t, 9, Part2(input))
}