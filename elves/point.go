package elves

type Point struct {
	X int
	Y int
}

func (p1 Point) Equals(p Point) bool {
	return p1.X == p.X && p1.Y == p.Y
}
