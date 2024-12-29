package q12

import (
	d "aoc2024/common/data"
	"aoc2024/inputs"
	"fmt"
	"iter"
	"strings"

	"github.com/stygian91/datastructs-go/set"
)

func Part1() {
	// content, err := inputs.GetInputFile("q12/demo.txt")
	content, err := inputs.GetInputFile("q12/main.txt")
	if err != nil {
		panic(err)
	}

	garden := parse(content)
	groups := group(&garden)
	answer := 0

	for _, group := range groups {
		perimeter := group.CountPerimeter(&garden)
		area := len(group.Items)
		answer += perimeter * area
	}

	fmt.Println("Part 1 answer:", answer)
}

type Group struct {
	Id    rune
	Items []d.Vec2i
}

func (this Group) String() string {
	return fmt.Sprintf("{%s %+v}", string(this.Id), this.Items)
}

func (this Group) CountPerimeter(garden *Gardens) int {
	perimeter := 0

	for _, pos := range this.Items {
		for _, neighbour := range garden.GetNeighbours(pos) {
			if this.Id != neighbour.Id {
				perimeter++
			}
		}

		if pos.X == 0 || pos.X == garden.Width-1 {
			perimeter++
		}

		if pos.Y == 0 || pos.Y == garden.Height-1 {
			perimeter++
		}
	}

	return perimeter
}

type Gardens struct {
	Grid   [][]rune
	Width  int
	Height int
}

type Neighbour struct {
	Pos d.Vec2i
	Id  rune
}

func (this Gardens) GetNeighbours(pos d.Vec2i) []Neighbour {
	neighbours := []Neighbour{}

	if pos.Y > 0 {
		neighbours = append(neighbours, Neighbour{
			Pos: d.Vec2i{Y: pos.Y - 1, X: pos.X},
			Id:  this.Grid[pos.Y-1][pos.X],
		})
	}

	if pos.X < this.Width-1 {
		neighbours = append(neighbours, Neighbour{
			Pos: d.Vec2i{Y: pos.Y, X: pos.X + 1},
			Id:  this.Grid[pos.Y][pos.X+1],
		})
	}

	if pos.Y < this.Height-1 {
		neighbours = append(neighbours, Neighbour{
			Pos: d.Vec2i{Y: pos.Y + 1, X: pos.X},
			Id:  this.Grid[pos.Y+1][pos.X],
		})
	}

	if pos.X > 0 {
		neighbours = append(neighbours, Neighbour{
			Pos: d.Vec2i{Y: pos.Y, X: pos.X - 1},
			Id:  this.Grid[pos.Y][pos.X-1],
		})
	}

	return neighbours
}

func parse(str string) Gardens {
	res := Gardens{}
	lines := strings.Split(strings.TrimSpace(str), "\n")
	res.Height = len(lines)

	for _, line := range lines {
		res.Width = len(line)
		row := []rune{}

		for _, r := range line {
			row = append(row, r)
		}

		res.Grid = append(res.Grid, row)
	}

	return res
}

func group(gardens *Gardens) []Group {
	ungrouped := set.New[d.Vec2i]()
	for y := 0; y < gardens.Height; y++ {
		for x := 0; x < gardens.Width; x++ {
			ungrouped.Add(d.Vec2i{X: x, Y: y})
		}
	}

	groups := []Group{}

	for ungrouped.Len() > 0 {
		plot, exists := getFirstSetEl(&ungrouped)
		if !exists {
			break
		}

		positions := []d.Vec2i{plot}
		ungrouped.Remove(plot)
		groups = append(groups, grow(gardens, &ungrouped, positions))
	}

	return groups
}

func grow(gardens *Gardens, ungrouped *set.Set[d.Vec2i], positions []d.Vec2i) Group {
	res := positions
	idPos := positions[0]
	id := gardens.Grid[idPos.Y][idPos.X]

	for {
		hasAdded := false

		for _, pos := range res {
			for _, neighbour := range gardens.GetNeighbours(pos) {
				if neighbour.Id != id || !ungrouped.Contains(neighbour.Pos) {
					continue
				}

				res = append(res, neighbour.Pos)
				ungrouped.Remove(neighbour.Pos)
				hasAdded = true
			}
		}

		if !hasAdded {
			break
		}
	}

	return Group{Items: res, Id: id}
}

func getFirstSetEl(s *set.Set[d.Vec2i]) (d.Vec2i, bool) {
	next, _ := iter.Pull(s.Seq())
	return next()
}
