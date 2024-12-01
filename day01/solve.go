package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parse(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse %s as uint64: %w", s, err))
	}
	return n
}

func Part1(input string) int {
	arr1 := []int64{}
	arr2 := []int64{}

	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, " ")
		arr1 = append(arr1, parse(parts[0]))
		arr2 = append(arr2, parse(parts[len(parts)-1]))
	}

	sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })
	sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })

	res := 0.0
	for i := 0; i < len(arr1); i++ {
		res += math.Abs(float64(arr1[i] - arr2[i]))
	}

	return int(res)
}

func Part2(input string) int {
	arr := []int64{}
	fr := map[int64]int{}

	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, " ")
		arr = append(arr, parse(parts[0]))

		v := parse(parts[len(parts)-1])
		fr[v]++
	}

	res := int(0)
	for i := 0; i < len(arr); i++ {
		res += int(arr[i]) * fr[arr[i]]
	}

	return res
}

func main() {
	b, _ := os.ReadFile("../data/day01/input.txt")
	input := string(b)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
