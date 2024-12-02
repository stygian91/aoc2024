package q02

import (
	"aoc2024/inputs"
	"fmt"
	"strings"
)

func Part1() {
	contents, err := inputs.GetInputFile("q02/main.txt")
	// contents, err := inputs.GetInputFile("q02/demo.txt")
	if err != nil {
		panic(err)
	}

	cnt := 0
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		nums, err := parseLine(line)
		if err != nil {
			panic(err)
		}

		if !isUnsafe(nums) {
			cnt++
		}
	}

	fmt.Printf("Part 1 answer: %d\n", cnt)
}
