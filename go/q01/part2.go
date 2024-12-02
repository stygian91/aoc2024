package q01

import (
	"aoc2024/inputs"
	"fmt"
	"strings"
)

func Part2() {
	dataMain, err := inputs.GetInputFile("q01/main.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(dataMain, "\n")
	lnums := []int{}
	rcnt := map[int]int{}

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		left, right, err := splitLine(line)
		if err != nil {
			panic(fmt.Errorf("Error on line %d: %w", i, err))
		}

		lnums = append(lnums, left)

		_, rexists := rcnt[right]
		if rexists {
			rcnt[right]++
		} else {
			rcnt[right] = 1
		}
	}

	sum := 0

	for _, lnum := range lnums {
		lrcnt, exists := rcnt[lnum]

		if exists {
			sum += lnum * lrcnt
		}
	}

	fmt.Printf("Part 2 answer: %d\n", sum)
}
