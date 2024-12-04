package q03

import (
	"aoc2024/inputs"
	"fmt"
	"strings"
)

func Part2() {
	// content, err := inputs.GetInputFile("q03/demo2.txt")
	content, err := inputs.GetInputFile("q03/main.txt")
	if err != nil {
		panic(err)
	}

	doIdx := allIndexes(content, "do()")
	dontIdx := allIndexes(content, "don't()")
	muls := parseMuls(content)
	sum := 0

	for _, m := range muls {
		if shouldDo(m.Pos, doIdx, dontIdx) {
			sum += m.A * m.B
		}
	}

	fmt.Printf("Part 2 answer: %d\n", sum)
}

func allIndexes(s, sub string) []int {
	res := []int{}
	rem := s[:]
	offset := 0

	for {
		idx := strings.Index(rem, sub)
		if idx == -1 {
			break
		}

		res = append(res, offset+idx)
		offset += idx+len(sub)
		rem = rem[idx+len(sub):]
	}

	return res
}

func shouldDo(idx int, doIdx, dontIdx []int) bool {
	maxDo := -1
	maxDont := -1

	for _, doI := range doIdx {
		if idx < doI {
			break
		}

		maxDo = doI
	}

	for _, dontI := range dontIdx {
		if idx < dontI {
			break
		}

		maxDont = dontI
	}

	if maxDo == -1 && maxDont == -1 {
		return true
	}

	return (maxDo > maxDont)
}
