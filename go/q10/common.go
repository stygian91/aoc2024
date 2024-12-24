package q10

import (
	d "aoc2024/common/data"
	"iter"
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
