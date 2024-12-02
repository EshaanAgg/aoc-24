package common

import (
	"fmt"
	"strconv"
)

func ParseInt(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse %s as uint64: %w", s, err))
	}

	return n
}
