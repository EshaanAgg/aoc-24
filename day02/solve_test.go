package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("../data/day02/sample.txt")
	input := string(b)
	fmt.Println(input)
	assert.Equal(t, Part1(input), int64(2))
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("../data/day02/sample.txt")
	input := string(b)

	assert.Equal(t, Part2(input), int64(4))
}
