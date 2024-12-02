package q02

import (
	"aoc2024/inputs"
	"fmt"
	"strconv"
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

func parseLine(line string) ([]int, error) {
	nums := []int{}

	for _, part := range strings.Split(line, " ") {
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nums, err
		}

		nums = append(nums, int(num))
	}

	return nums, nil
}

func isUnsafe(level []int) bool {
	if len(level) < 2 {
		return true
	}

	firstDiff := level[0] - level[1]
	if (firstDiff == 0 || absint(firstDiff) > 3) {
		return true
	}

	var isAsc bool
	if firstDiff < 0 {
		isAsc = true
	} else {
		isAsc = false
	}

	for i := 1; i < len(level) - 1; i++ {
		curr := level[i]
		next := level[i+1]
		diff := curr - next

		if (diff == 0) || (diff > 0 && isAsc) || (diff < 0 && !isAsc) || (absint(diff) > 3) {
			return true
		}
	}

	return false
}

func absint(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}
