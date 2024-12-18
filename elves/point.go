package elves

import (
	"fmt"
	"math"
)

var (
	DirectionUp        = Point{X: 0, Y: -1}
	DirectionDown      = Point{X: 0, Y: 1}
	DirectionLeft      = Point{X: -1, Y: 0}
	DirectionRight     = Point{X: 1, Y: 0}
	DirectionCardinals = []Point{DirectionUp, DirectionDown, DirectionLeft, DirectionRight}

	DirectionUpLeft    = Point{X: -1, Y: -1}
	DirectionUpRight   = Point{X: 1, Y: -1}
	DirectionDownLeft  = Point{X: -1, Y: 1}
	DirectionDownRight = Point{X: 1, Y: 1}
	DirectionDiagonals = []Point{DirectionUpLeft, DirectionUpRight, DirectionDownLeft, DirectionDownRight}

	DirectionAll = []Point{DirectionUp, DirectionDown, DirectionLeft, DirectionRight, DirectionUpLeft, DirectionUpRight, DirectionDownLeft, DirectionDownRight}
)

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

func (p1 Point) ShiftPos(p Point) Point {
	return Point{
		X: p1.X + p.X,
		Y: p1.Y + p.Y,
	}
}

func (p Point) Valid(maxX, maxY int) bool {
	return p.X >= 0 && p.Y >= 0 && p.X <= maxX && p.Y <= maxY
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) InBounds(p1, p2 Point) bool {
	return p.X >= p1.X && p.X <= p2.X && p.Y >= p1.Y && p.Y <= p2.Y
}

func (p Point) ManhattanDistance(p1 Point) int {
	return Abs(p.X-p1.X) + Abs(p.Y-p1.Y)
}

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}
