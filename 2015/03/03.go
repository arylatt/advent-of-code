package aoc201503

import "strconv"

type Point struct {
	X int
	Y int
}

func Part1(input string) (output string) {
	x, y := 0, 0
	houses := map[Point]int{
		{x, y}: 1,
	}

	for _, move := range input {
		switch move {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		}
		houses[Point{x, y}]++
	}

	return strconv.Itoa(len(houses))
}

func Part2(input string) (output string) {
	santas := [2]*Point{{0, 0}, {0, 0}}
	houses := map[Point]int{{0, 0}: 2}

	for i, move := range input {
		santa := santas[i%2]
		switch move {
		case '>':
			santa.X++
		case '<':
			santa.X--
		case '^':
			santa.Y++
		case 'v':
			santa.Y--
		}
		houses[*santa]++
	}

	return strconv.Itoa(len(houses))
}
