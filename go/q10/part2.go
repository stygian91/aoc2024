package q10

import (
	d "aoc2024/common/data"
	"aoc2024/inputs"
	"fmt"
)

func Part2() {
	// content, err := inputs.GetInputFile("q10/demo.txt")
	content, err := inputs.GetInputFile("q10/main.txt")
	if err != nil {
		panic(err)
	}

	grid := parse(content)
	sum := 0

	for _, head := range grid.Heads {
		s := []d.Vec2i{}
		scorePart2(&grid, head, &s)
		sum += len(s)
	}

	fmt.Println("Part 2 answer:", sum)
}

func scorePart2(grid *Grid, pos d.Vec2i, s *[]d.Vec2i) {
	value := grid.Values[pos.Y][pos.X]
	neighbours := filterNeighbours(value, grid.Neighbours(pos))
	if value == 8 {
		for nPos := range neighbours {
			*s = append(*s, nPos)
		}

		return
	}

	for nPos := range neighbours {
		scorePart2(grid, nPos, s)
	}
}
