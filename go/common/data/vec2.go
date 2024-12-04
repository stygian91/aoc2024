package data

type Vec2i struct {
	X, Y int
}

type Vec2f struct {
	X, Y float64
}

func (this Vec2i) Add(other Vec2i) Vec2i {
	return Vec2i{
		X: this.X + other.X,
		Y: this.Y + other.Y,
	}
}
