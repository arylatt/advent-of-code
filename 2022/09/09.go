package aoc202209

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func Part1(input string) (output string) {
	steps := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")
	head, tail := Point{0, 0}, Point{0, 0}
	tailVisited := map[Point]bool{tail: true}

	for _, step := range steps {
		instructions := strings.Split(step, " ")
		direction := instructions[0]
		count, _ := strconv.Atoi(instructions[1])

		for i := 0; i < count; i++ {
			switch direction {
			case "R":
				head.X++
			case "L":
				head.X--
			case "U":
				head.Y++
			case "D":
				head.Y--
			}

			if tail.X < head.X-1 || tail.X > head.X+1 || tail.Y < head.Y-1 || tail.Y > head.Y+1 {
				if tail.Y == head.Y {
					if direction == "R" {
						tail.X++
					} else {
						tail.X--
					}
				}

				if tail.X == head.X {
					if direction == "U" {
						tail.Y++
					} else {
						tail.Y--
					}
				}

				if tail.X < head.X && tail.Y < head.Y {
					if direction == "U" {
						tail = Point{head.X, head.Y - 1}
					} else {
						tail = Point{head.X - 1, head.Y}
					}
				}

				if tail.X > head.X && tail.Y > head.Y {
					if direction == "D" {
						tail = Point{head.X, head.Y + 1}
					} else {
						tail = Point{head.X + 1, head.Y}
					}
				}

				if tail.X > head.X && tail.Y < head.Y {
					if direction == "L" {
						tail = Point{head.X + 1, head.Y}
					} else {
						tail = Point{head.X, head.Y - 1}
					}

				}

				if tail.X < head.X && tail.Y > head.Y {
					if direction == "R" {
						tail = Point{head.X - 1, head.Y}
					} else {
						tail = Point{head.X, head.Y + 1}
					}
				}

				tailVisited[tail] = true
			}
		}
	}

	return strconv.Itoa(len(tailVisited))
}

func Part2(input string) (output string) {
	return
}
