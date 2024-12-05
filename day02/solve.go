package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eshaanagg/aoc24/common"
)

func parse(input string) [][]int {
	fin := [][]int{}

	for _, l := range strings.Split(input, "\n") {
		arr := []int{}
		for _, n := range strings.Split(l, " ") {
			arr = append(arr, common.ParseInt(n))
		}
		fin = append(fin, arr)
	}

	return fin
}

func check(a int, b int, incr bool) bool {
	if incr && a > b && a-b <= 3 {
		return true
	}

	if !incr && a < b && b-a <= 3 {
		return true
	}

	return false
}

func checkRow(arr []int) bool {
	n := len(arr)
	if n <= 1 {
		return true
	}

	if arr[0] == arr[1] {
		return false
	}

	incr := arr[0] < arr[1]
	for j := range n - 1 {
		if !check(arr[j+1], arr[j], incr) {
			return false
		}
	}

	return true
}

func checkRow2(arr []int) bool {
	if checkRow(arr) {
		return true
	}

	n := len(arr)
	for idx := range n {
		p := []int{}
		for j, ele := range arr {
			if j != idx {
				p = append(p, ele)
			}
		}
		if checkRow(p) {
			return true
		}
	}

	return false
}

func Part1(input string) int {
	rows := parse(input)
	cnt := 0

	for _, arr := range rows {
		if checkRow(arr) {
			cnt += 1
		}
	}

	return cnt
}

func Part2(input string) int {
	rows := parse(input)
	cnt := 0

	for _, arr := range rows {
		if checkRow2(arr) {
			cnt += 1
		}
	}

	return cnt
}

func main() {
	b, _ := os.ReadFile("./data/day02/input.txt")
	input := string(b)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
