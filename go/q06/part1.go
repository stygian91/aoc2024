package q06

import (
	"aoc2024/inputs"
	"fmt"
)

func Part1() {
	content, err := inputs.GetInputFile("q06/demo.txt")
	// content, err := inputs.GetInputFile("q06/main.txt")
	if err != nil {
		panic(err)
	}

	grid := parseGrid(content)
	fmt.Println(grid)

	fmt.Printf("Part 1 answer: \n")
}
