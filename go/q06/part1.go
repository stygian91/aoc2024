package q06

import (
	d "aoc2024/common/data"
	"aoc2024/inputs"
	"fmt"
)

func Part1() {
	// content, err := inputs.GetInputFile("q06/demo.txt")
	content, err := inputs.GetInputFile("q06/main.txt")
	if err != nil {
		panic(err)
	}

	grid := parseGrid(content)
	lines := []Line{}

	for {
		pos, inside := getNextPos(grid)
		if !inside {
			break
		}

		lines = append(lines, Line{
			Start: grid.Guard.Pos,
			End:   pos,
		})
		grid.Guard.Pos = pos
		grid.Guard.Dir = getNextDirection(grid.Guard.Dir)
	}

	var lastPos d.Vec2i
	switch grid.Guard.Dir {
	case North:
		lastPos = d.Vec2i{X: grid.Guard.Pos.X, Y: 0}
	case South:
		lastPos = d.Vec2i{X: grid.Guard.Pos.X, Y: grid.Height - 1}
	case East:
		lastPos = d.Vec2i{X: grid.Width - 1, Y: grid.Guard.Pos.Y}
	case West:
		lastPos = d.Vec2i{X: 0, Y: grid.Guard.Pos.Y}
	default:
		panic("Invalid direction")
	}

	lines = append(lines, Line{Start: grid.Guard.Pos, End: lastPos})
	answer := countUnique(lines)

	fmt.Printf("Part 1 answer: %d\n", answer)
}
