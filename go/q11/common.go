package q11

import (
	"strconv"
	"strings"
)

func parse(str string) []int {
	parts := strings.Split(strings.TrimSpace(str), " ")
	res := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		res = append(res, num)
	}

	return res
}
