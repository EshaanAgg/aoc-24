package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/eshaanagg/aoc24/common"
)

func Part1(input string) int {
	arr1 := []int{}
	arr2 := []int{}

	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, " ")
		arr1 = append(arr1, common.ParseInt(parts[0]))
		arr2 = append(arr2, common.ParseInt(parts[len(parts)-1]))
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	res := 0.0
	for i := 0; i < len(arr1); i++ {
		res += math.Abs(float64(arr1[i] - arr2[i]))
	}

	return int(res)
}

func Part2(input string) int {
	arr := []int{}
	fr := map[int]int{}

	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, " ")
		arr = append(arr, common.ParseInt(parts[0]))

		v := common.ParseInt(parts[len(parts)-1])
		fr[v]++
	}

	res := int(0)
	for i := 0; i < len(arr); i++ {
		res += int(arr[i]) * fr[arr[i]]
	}

	return res
}

func main() {
	b, _ := os.ReadFile("./data/day01/input.txt")
	input := string(b)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
