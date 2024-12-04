package q04

import (
	"aoc2024/inputs"
)

func Part1() {
	content, err := inputs.GetInputFile("q04/demo.txt")
	if err != nil {
		panic(err)
	}

	parseGrid(content)
}
