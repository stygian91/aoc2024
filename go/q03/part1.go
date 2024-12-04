package q03

import (
	"aoc2024/inputs"
	"fmt"
)

func Part1() {
	// content, err := inputs.GetInputFile("q03/demo.txt")
	content, err := inputs.GetInputFile("q03/main.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	muls := parseMuls(content)
	for _, m := range muls {
		sum += m.A * m.B
	}

	fmt.Printf("Part 1 answer: %d\n", sum)
}
