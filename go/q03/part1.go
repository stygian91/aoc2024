package q03

import (
	"aoc2024/inputs"
	// "fmt"
	"strings"
)

func Part1() {
	content, err := inputs.GetInputFile("q03/demo.txt")
	if err != nil {
		panic(err)
	}

	remaining := content[:]
	for len(remaining) > 0 {
		idx := strings.Index(remaining, "mul(")
		if idx == -1 {
			break
		}
	}
}
