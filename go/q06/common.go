package q06

import (
	d "aoc2024/common/data"
	"strings"
)

type Dir uint8

const (
	North = iota
	East
	South
	West
)

type Guard struct {
	Pos d.Vec2i
	Dir Dir
}

type Grid struct {
	Guard        Guard
	Obstructions []d.Vec2i
}

func parseGrid(str string) Grid {
	lines := strings.Split(strings.TrimSpace(str), "\n")
	res := Grid{}

	for y, line := range lines {
		for x, r := range line {
			switch r {
			case '^':
				res.Guard = Guard{Pos: d.Vec2i{X: x, Y: y}, Dir: North}
			case '#':
				res.Obstructions = append(res.Obstructions, d.Vec2i{X: x, Y: y})
			}
		}
	}

	return res
}
