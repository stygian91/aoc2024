package q06_test

import (
	"aoc2024/common/data"
	"aoc2024/inputs"
	"aoc2024/q06"

	"slices"
	"testing"
)

var grid q06.Grid

func init() {
	content, err := inputs.GetInputFile("q06/demo.txt")
	if err != nil {
		panic(err)
	}

	grid = q06.ParseGrid(content)
}

func TestIsInfiniteFalseNegative(t *testing.T) {
	placements := []data.Vec2i{
		{X: 3, Y: 6},
		{X: 6, Y: 7},
		{X: 7, Y: 7},
		{X: 1, Y: 8},
		{X: 3, Y: 8},
		{X: 7, Y: 9},
	}

	for _, placement := range placements {
		gcopy := q06.CloneGrid(grid)
		gcopy.AppendRow(placement.Y, placement.X)
		gcopy.AppendCol(placement.X, placement.Y)

		if !q06.IsInfinite(gcopy) {
			t.Errorf("Expected placement `%+v` to be infinite\n", placement)
		}
	}
}

func TestIsInfiniteFalsePositive(t *testing.T) {
	placements := []data.Vec2i{
		{X: 3, Y: 6},
		{X: 6, Y: 7},
		{X: 7, Y: 7},
		{X: 1, Y: 8},
		{X: 3, Y: 8},
		{X: 7, Y: 9},
	}

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			pos := data.Vec2i{X: x, Y: y}
			row := grid.Rows[y]
			_, e := row.Search(x)
			if e || (grid.Guard.Pos == pos) {
				continue
			}

			gcopy := q06.CloneGrid(grid)
			gcopy.AppendRow(y, x)
			gcopy.AppendCol(x, y)

			if q06.IsInfinite(gcopy) && !slices.Contains(placements, pos) {
				t.Errorf("Pos %+v is a false positive\n", pos)
			}
		}
	}
}
