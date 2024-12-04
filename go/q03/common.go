package q03

import (
	"strconv"
	"strings"
)

type mul struct {
	A, B, Pos int
}

func parseNumber(str string) (string, bool) {
	idx := 0

	returnRes := func() (string, bool) {
		res := str[0:idx]
		if len(res) == 0 {
			return "", false
		}

		return res, true
	}

	for i, r := range str {
		idx = i
		if r < '0' || r > '9' {
			return returnRes()
		}
	}

	return returnRes()
}

func parseMuls(str string) []mul {
	res := []mul{}
	remaining := str[:]
	offset := 0
	startOffset := 0

	eat := func(amount int) {
		remaining = remaining[amount:]
		offset += amount
	}

	for len(remaining) > 0 {
		idx := strings.Index(remaining, "mul(")
		if idx == -1 {
			break
		}

		startOffset = idx + offset
		eat(idx + 4)
		val, found := parseNumber(remaining)
		if !found {
			continue
		}

		eat(len(val))
		a, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			continue
		}

		if remaining[0] != ',' {
			continue
		}

		eat(1)
		val, found = parseNumber(remaining)
		if !found {
			continue
		}

		eat(len(val))
		b, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			continue
		}

		if remaining[0] != ')' {
			continue
		}

		res = append(res, mul{
			A:   int(a),
			B:   int(b),
			Pos: startOffset,
		})
	}

	return res
}
