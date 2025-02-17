package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("../data/day02/sample.txt")
	input := string(b)

	assert.Equal(t, Part1(input), 2)
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("../data/day02/sample.txt")
	input := string(b)

	assert.Equal(t, Part2(input), 4)
}
