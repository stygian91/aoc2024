package q02

import (
	"strconv"
	"strings"
)

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

func absint(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

func isUnsafe(level []int) bool {
	if len(level) < 2 {
		return true
	}

	firstDiff := level[0] - level[1]
	if firstDiff == 0 || absint(firstDiff) > 3 {
		return true
	}

	var isAsc bool
	if firstDiff < 0 {
		isAsc = true
	} else {
		isAsc = false
	}

	for i := 1; i < len(level)-1; i++ {
		curr := level[i]
		next := level[i+1]
		diff := curr - next

		if (diff == 0) || (diff > 0 && isAsc) || (diff < 0 && !isAsc) || (absint(diff) > 3) {
			return true
		}
	}

	return false
}
