package q04

import (
	d "aoc2024/common/data"
	"aoc2024/inputs"
	"fmt"
)

func Part2() {
	// content, err := inputs.GetInputFile("q04/demo.txt")
	content, err := inputs.GetInputFile("q04/main.txt")
	if err != nil {
		panic(err)
	}

	grid := parseGrid(content)
	cnt := 0

	for y, row := range grid.Cells {
		for x, cell := range row {
			if cell != 'A' {
				continue
			}

			nw := d.Vec2i{X: x - 1, Y: y - 1}
			ne := d.Vec2i{X: x + 1, Y: y - 1}
			nwStr := string(grid.Walk(nw, d.Vec2i{X: 1, Y: 1}, 3))
			neStr := string(grid.Walk(ne, d.Vec2i{X: -1, Y: 1}, 3))
			nwMatches := nwStr == "MAS" || nwStr == "SAM"
			neMatches := neStr == "MAS" || neStr == "SAM"

			if nwMatches && neMatches {
				cnt++
			}
		}
	}

	fmt.Printf("Part 2 answer: %d\n", cnt)
}
