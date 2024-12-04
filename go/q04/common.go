package q04

import (
	"aoc2024/common/data"
	"strings"
)

type Grid struct {
	Cells [][]rune
}

func (this Grid) IsInside(_ data.Vec2i) bool {
	// TODO:
	return false
}

func (this Grid) At(pos data.Vec2i) rune {
	return this.Cells[pos.Y][pos.X]
}

func parseGrid(str string) Grid {
	grid := Grid{}
	cells := [][]rune{}
	row := []rune{}

	for _, r := range strings.TrimSpace(str) {
		if r == '\r' {
			continue
		}

		if r == '\n' {
			cells = append(cells, row)
			row = []rune{}
			continue
		}

		row = append(row, r)
	}

	grid.Cells = cells
	return grid
}

func walk(grid Grid, start, dir data.Vec2i) []rune {
	res := []rune{}
	next := start.Add(dir)

	for {
		if !grid.IsInside(next) {
			break
		}

		res = append(res, grid.At(next))
		next = next.Add(dir)
	}

	return res
}

