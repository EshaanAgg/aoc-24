package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/eshaanagg/aoc24/common"
)

func parse(input string) (map[int][]int, [][]int) {
	lines := strings.Split(input, "\n")

	after := make(map[int][]int)
	ord := [][]int{}

	ruleParsing := true

	for _, l := range lines {
		if l == "" {
			ruleParsing = false
			continue
		}

		if ruleParsing {
			parts := strings.Split(l, "|")
			u := common.ParseInt(parts[0])
			v := common.ParseInt(parts[1])

			after[u] = append(after[u], v)
		} else {
			parts := strings.Split(l, ",")

			arr := []int{}
			for _, p := range parts {
				arr = append(arr, common.ParseInt(p))
			}

			ord = append(ord, arr)
		}
	}

	return after, ord
}

func checkValid(after map[int][]int, arr []int) bool {
	seen := make(map[int]bool)

	for _, ele := range arr {
		for _, nxtEle := range after[ele] {
			if seen[nxtEle] {
				return false
			}
		}
		seen[ele] = true
	}

	return true
}

func makeValid(after map[int][]int, arr []int) {
	slices.SortStableFunc(arr, func(i, j int) int {
		for _, nxtEle := range after[i] {
			if nxtEle == j {
				return 1
			}
		}

		for _, nxtEle := range after[j] {
			if nxtEle == i {
				return -1
			}
		}

		return 0
	})
}

func Part1(input string) int {
	after, ord := parse(input)
	res := 0

	for _, arr := range ord {
		if checkValid(after, arr) {
			n := len(arr)
			res += arr[n/2]
		}
	}

	return res
}

func Part2(input string) int {
	after, ord := parse(input)
	res := 0

	for _, arr := range ord {
		if !checkValid(after, arr) {
			makeValid(after, arr)
			n := len(arr)
			res += arr[n/2]
		}
	}

	return res
}

func main() {
	b, _ := os.ReadFile("./data/day05/input.txt")
	input := string(b)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
