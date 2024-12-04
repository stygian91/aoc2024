package q04

import (
	d "aoc2024/common/data"
	"aoc2024/inputs"
	"fmt"
)

func Part1() {
	// content, err := inputs.GetInputFile("q04/demo.txt")
	content, err := inputs.GetInputFile("q04/main.txt")
	if err != nil {
		panic(err)
	}

	grid := parseGrid(content)
	cnt := 0
	dirs := []d.Vec2i{
		{X: 1, Y: 0},   // E
		{X: 1, Y: 1},   // SE
		{X: 0, Y: 1},   // S
		{X: -1, Y: 1},  // SW
		{X: -1, Y: 0},  // W
		{X: -1, Y: -1}, // NW
		{X: 0, Y: -1},  // N
		{X: 1, Y: -1},  // NE
	}

	for y, row := range grid.Cells {
		for x, cell := range row {
			if cell != 'X' {
				continue
			}

			for _, dir := range dirs {
				letters := grid.Walk(d.Vec2i{X: x, Y: y}, dir, 4)
				if string(letters) == "XMAS" {
					cnt++
				}
			}
		}
	}

	fmt.Printf("Part 1 answer: %d\n", cnt)
}
