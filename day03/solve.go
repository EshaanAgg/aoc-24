package main

import (
	"fmt"
	"os"
)

type Solver struct {
	s       string
	idx     int
	n       int
	enabled bool
}

func (so *Solver) isDigit(idx int) (bool, int) {
	if idx >= so.n {
		return false, 0
	}
	return so.s[idx] >= '0' && so.s[idx] <= '9', int(so.s[idx] - '0')

}

func (so *Solver) parseNum(idx int) (int, int) {
	b, v := so.isDigit(idx)
	if !b {
		return 0, 0
	}

	len := 1
	for {
		b, nv := so.isDigit(idx + len)
		if !b {
			break
		}
		len += 1
		v = 10*v + nv
	}

	return len, v
}

func (so *Solver) changeState() {
	idx := so.idx

	if idx+4 <= so.n && so.s[idx:idx+4] == "do()" {
		so.enabled = true
		so.idx += 4
	} else if idx+7 <= so.n && so.s[idx:idx+7] == "don't()" {
		so.enabled = false
		so.idx += 7
	}
}

func (so *Solver) parseInstruction(idx int) int {
	if idx+4 > so.n || so.s[idx:idx+4] != "mul(" {
		so.idx += 1
		return 0
	}
	idx += 4

	l1, d1 := so.parseNum(idx)
	if l1 < 1 || l1 > 3 {
		so.idx += 1
		return 0
	}
	idx += l1

	if idx >= so.n || so.s[idx] != ',' {
		so.idx += 1
		return 0
	}
	idx += 1

	l2, d2 := so.parseNum(idx)
	if l2 < 1 || l2 > 3 {
		so.idx += 1
		return 0
	}
	idx += l2

	if idx >= so.n || so.s[idx] != ')' {
		so.idx += 1
		return 0
	}
	idx += 1

	so.idx = idx
	return d1 * d2
}

func Part1(input string) int {
	so := Solver{
		s:       input,
		idx:     0,
		n:       len(input),
		enabled: true,
	}

	res := 0
	for so.idx < so.n {
		r := so.parseInstruction(so.idx)
		res += r
	}

	return res
}

func Part2(input string) int {
	so := Solver{
		s:       input,
		idx:     0,
		n:       len(input),
		enabled: true,
	}

	res := 0
	for so.idx < so.n {
		so.changeState()
		r := so.parseInstruction(so.idx)

		if so.enabled {
			res += r
		}
	}

	return res
}

func main() {
	b, _ := os.ReadFile("./data/day03/input.txt")
	input := string(b)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
