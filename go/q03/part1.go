package q03

import (
	"aoc2024/inputs"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	// content, err := inputs.GetInputFile("q03/demo.txt")
	content, err := inputs.GetInputFile("q03/main.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	remaining := content[:]
	for len(remaining) > 0 {
		idx := strings.Index(remaining, "mul(")
		if idx == -1 {
			break
		}

		remaining = remaining[idx+4:]
		val, found := parseNumber(remaining)
		if !found {
			continue
		}

		remaining = remaining[len(val):]
		a, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			continue
		}

		if remaining[0] != ',' {
			continue
		}

		remaining = remaining[1:]
		val, found = parseNumber(remaining)
		if !found {
			continue
		}

		remaining = remaining[len(val):]
		b, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			continue
		}

		if remaining[0] != ')' {
			continue
		}

		fmt.Printf("a: %d; b: %d\n", a, b)
		sum += int(a) * int(b)
	}

	fmt.Printf("Part 1 answer: %d\n", sum)
}
