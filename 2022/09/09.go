package aoc202209

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func Move(length int, steps []string) (tailVisited map[Point]bool) {
	rope := []Point{}
	for i := 0; i < length; i++ {
		rope = append(rope, Point{0, 0})
	}

	tailVisited = map[Point]bool{rope[length-1]: true}

	for _, step := range steps {
		instructions := strings.Split(step, " ")
		direction := instructions[0]
		count, _ := strconv.Atoi(instructions[1])

		for i := 0; i < count; i++ {
			previousHead := rope[0]

			switch direction {
			case "R":
				rope[0].X++
			case "L":
				rope[0].X--
			case "U":
				rope[0].Y++
			case "D":
				rope[0].Y--
			}

			for headI := 0; headI < length-1; headI++ {
				previousTail := rope[headI+1]

				if rope[headI+1].X < rope[headI].X-1 || rope[headI+1].X > rope[headI].X+1 || rope[headI+1].Y < rope[headI].Y-1 || rope[headI+1].Y > rope[headI].Y+1 {
					headDiff := Point{rope[headI].X - previousHead.X, rope[headI].Y - previousHead.Y}

					if rope[headI].X-rope[headI+1].X == 0 || rope[headI].Y-rope[headI+1].Y == 0 {
						coordDiff := Point{rope[headI].X - rope[headI+1].X, rope[headI].Y - rope[headI+1].Y}
						if coordDiff.X != 0 {
							coordDiff.X = coordDiff.X / int(math.Abs(float64(coordDiff.X)))
						} else {
							coordDiff.Y = coordDiff.Y / int(math.Abs(float64(coordDiff.Y)))
						}
						rope[headI+1] = Point{rope[headI+1].X + coordDiff.X, rope[headI+1].Y + coordDiff.Y}
					} else if headDiff.X != 0 && headDiff.Y != 0 && !(headDiff.X == 0 && headDiff.Y == 0) {
						rope[headI+1] = Point{rope[headI+1].X + headDiff.X, rope[headI+1].Y + headDiff.Y}
					} else {
						rope[headI+1] = previousHead
					}

					previousHead = previousTail
				} else {
					break
				}
			}

			tailVisited[rope[length-1]] = true
		}
	}

	return
}

func Part1(input string) (output string) {
	steps := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")

	return strconv.Itoa(len(Move(2, steps)))
}

func Part2(input string) (output string) {
	steps := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")

	return strconv.Itoa(len(Move(10, steps)))
}
