package elves

type Point struct {
	X int
	Y int
}

func (p1 Point) Equals(p Point) bool {
	return p1.X == p.X && p1.Y == p.Y
}

func (p1 Point) Shift(x, y int) Point {
	return Point{
		X: p1.X + x,
		Y: p1.Y + y,
	}
}

func (p Point) Valid(maxX, maxY int) bool {
	return p.X >= 0 && p.Y >= 0 && p.X <= maxX && p.Y <= maxY
}
