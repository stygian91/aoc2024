package q06

import (
	d "aoc2024/common/data"
	"slices"
	"strings"
)

type Dir uint8

func (this Dir) String() string {
	switch this {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	default:
		return "Unknown"
	}
}

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
	Guard Guard
	Rows  map[int][]int
	Cols  map[int][]int

	Width, Height int
}

func parseGrid(str string) Grid {
	lines := strings.Split(strings.TrimSpace(str), "\n")
	res := Grid{
		Rows: map[int][]int{},
		Cols: map[int][]int{},
	}

	res.Height = len(lines)
	for y, line := range lines {
		res.Width = len(line)

		for x, r := range line {
			switch r {
			case '^':
				res.Guard = Guard{Pos: d.Vec2i{X: x, Y: y}, Dir: North}
			case '#':
				res.Rows[y] = append(res.Rows[y], x)
				res.Cols[x] = append(res.Cols[x], y)
			}
		}
	}

	for y := range res.Rows {
		slices.Sort(res.Rows[y])
	}

	for x := range res.Cols {
		slices.Sort(res.Cols[x])
	}

	return res
}

func findClosestNorth(grid Grid) int {
	closest := -1
	x := grid.Guard.Pos.X

	for i := 0; i < len(grid.Cols[x]); i++ {
		y := grid.Cols[x][i]

		if y >= grid.Guard.Pos.Y {
			break
		}

		closest = y
	}

	return closest
}

func findClosestSouth(grid Grid) int {
	closest := -1
	x := grid.Guard.Pos.X

	for i := len(grid.Cols[x]) - 1; i >= 0; i-- {
		y := grid.Cols[x][i]

		if y <= grid.Guard.Pos.Y {
			break
		}

		closest = y
	}

	return closest
}

func findClosestEast(grid Grid) int {
	closest := -1
	y := grid.Guard.Pos.Y

	for i := len(grid.Rows[y]) - 1; i >= 0; i-- {
		x := grid.Rows[y][i]

		if x <= grid.Guard.Pos.X {
			break
		}

		closest = x
	}

	return closest
}

func findClosestWest(grid Grid) int {
	closest := -1
	y := grid.Guard.Pos.Y

	for i := 0; i < len(grid.Rows[y]); i++ {
		x := grid.Rows[y][i]

		if x >= grid.Guard.Pos.X {
			break
		}

		closest = x
	}

	return closest
}

func getNextDirection(dir Dir) Dir {
	switch dir {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		panic("Unexpected direction")
	}
}

func getNextPos(grid Grid) (d.Vec2i, bool) {
	switch grid.Guard.Dir {
	case North:
		y := findClosestNorth(grid)
		if y == -1 {
			return d.Vec2i{}, false
		}
		return d.Vec2i{X: grid.Guard.Pos.X, Y: y + 1}, true
	case South:
		y := findClosestSouth(grid)
		if y == -1 {
			return d.Vec2i{}, false
		}
		return d.Vec2i{X: grid.Guard.Pos.X, Y: y - 1}, true
	case East:
		x := findClosestEast(grid)
		if x == -1 {
			return d.Vec2i{}, false
		}
		return d.Vec2i{X: x - 1, Y: grid.Guard.Pos.Y}, true
	case West:
		x := findClosestWest(grid)
		if x == -1 {
			return d.Vec2i{}, false
		}
		return d.Vec2i{X: x + 1, Y: grid.Guard.Pos.Y}, true
	default:
		panic("Unexpected direction")
	}
}

func countUnique(lines []Line) int {
	type empty struct {}
	visited := map[d.Vec2i]empty{}

	walkH := func (line Line) {
		var start, end int
		if line.Start.X < line.End.X {
			start = line.Start.X
			end = line.End.X
		} else {
			start = line.End.X
			end = line.Start.X
		}

		for i := start; i <= end; i++ {
			idx := d.Vec2i{ X: i, Y: line.Start.Y }
			visited[idx] = empty{}
		}
	}

	walkV := func (line Line) {
		var start, end int
		if line.Start.Y < line.End.Y {
			start = line.Start.Y
			end = line.End.Y
		} else {
			start = line.End.Y
			end = line.Start.Y
		}

		for i := start; i <= end; i++ {
			idx := d.Vec2i{ Y: i, X: line.Start.X }
			visited[idx] = empty{}
		}
	}

	for _, line := range lines {
		if line.IsHorizontal() {
			walkH(line)
		} else {
			walkV(line)
		}
	}

	return len(visited)
}

type Line struct {
	Start, End d.Vec2i
}

func (this Line) IsHorizontal() bool {
	return this.Start.Y == this.End.Y
}
