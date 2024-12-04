package main

import (
	"fmt"
	"os"
	"strings"
)

type Checker struct {
	grid []string
	Y    int
	X    int
}

func NewChecker(s string) Checker {
	grid := strings.Split(s, "\n")
	return Checker{
		grid: grid,
		Y:    len(grid),
		X:    len(grid[0]),
	}
}

func (ch *Checker) checkDir(x int, y int, dX int, dY int, search string) bool {
	for idx := range len(search) {
		if x < 0 || y < 0 || x >= ch.X || y >= ch.Y || ch.grid[y][x] != search[idx] {
			return false
		}
		x += dX
		y += dY
	}

	return true
}

func (ch *Checker) checkAll(x int, y int, search string) int {
	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	cnt := 0
	for _, dir := range dirs {
		if ch.checkDir(x, y, dir[0], dir[1], search) {
			cnt += 1
		}
	}
	return cnt
}

func Part1(input string) int {
	ch := NewChecker(input)
	cnt := 0

	for y := range ch.Y {
		for x := range ch.X {
			cnt += ch.checkAll(x, y, "XMAS")
		}
	}

	return cnt
}

func Part2(input string) int {
	ch := NewChecker(input)
	cnt := 0

	for y := range ch.Y - 2 {
		for x := range ch.X - 2 {
			b1 := ch.checkDir(x, y, 1, 1, "MAS") || ch.checkDir(x, y, 1, 1, "SAM")
			b2 := ch.checkDir(x+2, y, -1, 1, "MAS") || ch.checkDir(x+2, y, -1, 1, "SAM")
			if b1 && b2 {
				cnt += 1
			}
		}
	}

	return cnt
}

func main() {
	b, _ := os.ReadFile("./data/day04/input.txt")
	input := string(b)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
