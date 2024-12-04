package q04

import (
	"aoc2024/common/data"
	"strings"
)

type Grid struct {
	Cells [][]rune
	W, H  int
}

func (this Grid) IsInside(pos data.Vec2i) bool {
	return pos.X >= 0 && pos.X < this.W && pos.Y >= 0 && pos.Y < this.H
}

func (this Grid) At(pos data.Vec2i) rune {
	return this.Cells[pos.Y][pos.X]
}

func (this Grid) Walk(start, dir data.Vec2i, count int) []rune {
	res := []rune{}
	next := start

	for i := 0; i < count; i++ {
		if !this.IsInside(next) {
			break
		}

		res = append(res, this.At(next))
		next = next.Add(dir)
	}

	return res
}

func parseGrid(str string) Grid {
	grid := Grid{}
	cells := [][]rune{}
	row := []rune{}
	s := strings.TrimSpace(str)

	for i, r := range s {
		if r == '\r' {
			continue
		}

		if r == '\n' {
			cells = append(cells, row)
			row = []rune{}
			continue
		}

		row = append(row, r)

		if i == len(s)-1 {
			cells = append(cells, row)
		}
	}

	grid.Cells = cells
	grid.H = len(cells)
	grid.W = len(cells[0])
	return grid
}
