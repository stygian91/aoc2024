package q06

import (
	d "aoc2024/common/data"
	"slices"
	"strings"

	"github.com/stygian91/datastructs-go/bst"
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
	Rows  map[int]bst.BST[int, struct{}]
	Cols  map[int]bst.BST[int, struct{}]

	Width, Height int
}

func ParseGrid(str string) Grid {
	lines := strings.Split(strings.TrimSpace(str), "\n")
	res := Grid{
		Rows: make(map[int]bst.BST[int, struct{}]),
		Cols: make(map[int]bst.BST[int, struct{}]),
	}

	res.Height = len(lines)
	for y, line := range lines {
		res.Width = len(line)

		for x, r := range line {
			switch r {
			case '^':
				res.Guard = Guard{Pos: d.Vec2i{X: x, Y: y}, Dir: North}
			case '#':
				res.AppendRow(y, x)
				res.AppendCol(x, y)
			}
		}
	}

	for y := range res.Rows {
		res.Rows[y] = res.Rows[y].NewBalanced()
	}

	for x := range res.Cols {
		res.Cols[x] = res.Cols[x].NewBalanced()
	}

	return res
}

func FindClosestNorth(grid Grid) int {
	closest := -1
	x := grid.Guard.Pos.X
	col, e := grid.Cols[x]
	if !e {
		return -1
	}

	for v := range col.InOrderSeq() {
		if v.Value >= grid.Guard.Pos.Y {
			break
		}

		closest = v.Value
	}

	return closest
}

func findClosestSouth(grid Grid) int {
	closest := -1
	x := grid.Guard.Pos.X
	col, e := grid.Cols[x]
	if !e {
		return -1
	}

	for v := range col.PostOrderSeq() {
		if v.Value <= grid.Guard.Pos.Y {
			break
		}

		closest = v.Value
	}

	return closest
}

func findClosestEast(grid Grid) int {
	closest := -1
	y := grid.Guard.Pos.Y
	row, e := grid.Rows[y]
	if !e {
		return -1
	}

	for v := range row.PostOrderSeq() {
		if v.Value <= grid.Guard.Pos.X {
			break
		}

		closest = v.Value
	}

	return closest
}

func findClosestWest(grid Grid) int {
	closest := -1
	y := grid.Guard.Pos.Y
	row, e := grid.Rows[y]
	if !e {
		return -1
	}

	for v := range row.InOrderSeq() {
		if v.Value >= grid.Guard.Pos.X {
			break
		}

		closest = v.Value
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
		y := FindClosestNorth(grid)
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
	type empty struct{}
	visited := map[d.Vec2i]empty{}

	walkH := func(line Line) {
		var start, end int
		if line.Start.X < line.End.X {
			start = line.Start.X
			end = line.End.X
		} else {
			start = line.End.X
			end = line.Start.X
		}

		for i := start; i <= end; i++ {
			idx := d.Vec2i{X: i, Y: line.Start.Y}
			visited[idx] = empty{}
		}
	}

	walkV := func(line Line) {
		var start, end int
		if line.Start.Y < line.End.Y {
			start = line.Start.Y
			end = line.End.Y
		} else {
			start = line.End.Y
			end = line.Start.Y
		}

		for i := start; i <= end; i++ {
			idx := d.Vec2i{Y: i, X: line.Start.X}
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

	Dir Dir
}

func (this Line) IsHorizontal() bool {
	return this.Start.Y == this.End.Y
}

func IsInfinite(grid Grid) bool {
	lines := []Line{}

	update := func(pos d.Vec2i) {
		lines = append(lines, Line{
			Start: grid.Guard.Pos,
			End:   pos,
			Dir:   grid.Guard.Dir,
		})
		grid.Guard.Pos = pos
		grid.Guard.Dir = getNextDirection(grid.Guard.Dir)
	}

	for {
		pos, inside := getNextPos(grid)
		if !inside {
			return false
		}

		sIdx := IsVisited(lines, pos)
		if sIdx != -1 && (lines[sIdx].Dir == getNextDirection(grid.Guard.Dir)) {
			return true
		}

		update(pos)
	}
}

func IsVisited(lines []Line, pos d.Vec2i) int {
	return slices.IndexFunc(lines, func(line Line) bool {
		return line.Start == pos
	})
}

func CloneGrid(grid Grid) Grid {
	res := Grid{}

	clone := func(mapOfInts map[int]bst.BST[int, struct{}]) map[int]bst.BST[int, struct{}] {
		mapRes := map[int]bst.BST[int, struct{}]{}

		for k := range mapOfInts {
			mapRes[k] = mapOfInts[k].NewBalanced()
		}

		return mapRes
	}

	res.Width = grid.Width
	res.Height = grid.Height
	res.Guard = Guard{
		Pos: d.Vec2i{X: grid.Guard.Pos.X, Y: grid.Guard.Pos.Y},
		Dir: grid.Guard.Dir,
	}

	res.Rows = clone(grid.Rows)
	res.Cols = clone(grid.Cols)

	return res
}

func (this *Grid) AppendRow(y, x int) {
	row, e := this.Rows[y]
	if !e {
		row = bst.NewBST(x, struct{}{})
		this.Rows[y] = row
		return
	}

	row.Add(x, struct{}{})
	this.Rows[y] = row
}

func (this *Grid) AppendCol(x, y int) {
	col, e := this.Cols[x]
	if !e {
		col = bst.NewBST(y, struct{}{})
		this.Cols[x] = col
		return
	}

	col.Add(y, struct{}{})
	this.Cols[x] = col
}
