package q02

import (
	"aoc2024/inputs"
	"fmt"
	"slices"
	"strings"
)

func Part2() {
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

		unsafe := isUnsafe(nums)
		if unsafe {
			unsafe = tryRemovingOne(nums)
		}

		if !unsafe {
			cnt++
		}
	}

	fmt.Printf("Part 2 answer: %d\n", cnt)
}

func tryRemovingOne(level []int) bool {
	for i := 0; i < len(level); i++ {
		spliced := slices.Concat(level[0:i], level[i+1:])
		if !isUnsafe(spliced) {
			return false
		}
	}

	return true
}
