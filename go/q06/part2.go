package q06

import (
	d "aoc2024/common/data"
	"aoc2024/inputs"
	"fmt"
)

func Part2() {
	// content, err := inputs.GetInputFile("q06/demo.txt")
	content, err := inputs.GetInputFile("q06/main.txt")
	if err != nil {
		panic(err)
	}

	grid := ParseGrid(content)
	cnt := 0

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			row := grid.Rows[y]
			_, e := row.Search(x)
			if e || (grid.Guard.Pos == d.Vec2i{X: x, Y: y}) {
				continue
			}

			gcopy := CloneGrid(grid)
			gcopy.AppendRow(y, x)
			gcopy.AppendCol(x, y)

			if IsInfinite(gcopy) {
				cnt++
			}
		}
	}

	fmt.Printf("Part 1 answer: %d\n", cnt)
}
