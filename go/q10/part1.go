package q10

import (
	d "aoc2024/common/data"
	"aoc2024/inputs"
	"fmt"
	"iter"
	"strconv"
	"strings"

	"github.com/stygian91/datastructs-go/set"
)

func Part1() {
	// content, err := inputs.GetInputFile("q10/demo.txt")
	content, err := inputs.GetInputFile("q10/main.txt")
	if err != nil {
		panic(err)
	}

	grid := parse(content)
	sum := 0

	for _, head := range grid.Heads {
		s := set.New[d.Vec2i]()
		scorePart1(&grid, head, &s)
		sum += s.Len()
	}

	fmt.Println("Part 1 answer:", sum)
}

func scorePart1(grid *Grid, pos d.Vec2i, s *set.Set[d.Vec2i]) {
	value := grid.Values[pos.Y][pos.X]
	neighbours := filterNeighbours(value, grid.Neighbours(pos))
	if value == 8 {
		for nPos := range neighbours {
			s.Add(nPos)
		}

		return
	}

	for nPos := range neighbours {
		scorePart1(grid, nPos, s)
	}
}

func filterNeighbours(currValue uint8, neighbours iter.Seq2[d.Vec2i, uint8]) iter.Seq2[d.Vec2i, uint8] {
	return func(yield func(d.Vec2i, uint8) bool) {
		for pos, neighbour := range neighbours {
			if neighbour != currValue+1 {
				continue
			}

			if !yield(pos, neighbour) {
				return
			}
		}
	}
}

func parse(str string) Grid {
	res := Grid{Values: [][]uint8{}}
	lines := strings.Split(strings.TrimSpace(str), "\n")
	res.Height = len(lines)

	for y, line := range lines {
		row := []uint8{}
		res.Width = len(line)

		for x, r := range line {
			digit, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			if digit == 0 {
				res.Heads = append(res.Heads, d.Vec2i{X: x, Y: y})
			}

			row = append(row, uint8(digit))
		}

		res.Values = append(res.Values, row)
	}

	return res
}
