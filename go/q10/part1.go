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

type Grid struct {
	Values [][]uint8
	Heads  []d.Vec2i

	Width, Height int
}

type Dir uint8

const (
	North Dir = iota
	East
	South
	West
)

func (this Grid) Neighbours(pos d.Vec2i) iter.Seq2[d.Vec2i, uint8] {
	return func(yield func(d.Vec2i, uint8) bool) {
		for dir := North; dir <= West; dir++ {
			x, y := -1, -1

			switch dir {
			case North:
				if pos.Y > 0 {
					x = pos.X
					y = pos.Y - 1
				}
			case East:
				if pos.X < this.Width-1 {
					x = pos.X + 1
					y = pos.Y
				}
			case South:
				if pos.Y < this.Height-1 {
					x = pos.X
					y = pos.Y + 1
				}
			case West:
				if pos.X > 0 {
					x = pos.X - 1
					y = pos.Y
				}
			}

			if x == -1 {
				continue
			}

			if !yield(d.Vec2i{X: x, Y: y}, this.Values[y][x]) {
				return
			}
		}
	}
}

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
		score(&grid, head, &s)
		sum += s.Len()
	}

	fmt.Println("Part 2 answer:", sum)
}

func score(grid *Grid, pos d.Vec2i, s *set.Set[d.Vec2i]) {
	value := grid.Values[pos.Y][pos.X]
	neighbours := filterNeighbours(value, grid.Neighbours(pos))
	if value == 8 {
		for nPos := range neighbours {
			s.Add(nPos)
		}

		return
	}

	for nPos := range neighbours {
		score(grid, nPos, s)
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
