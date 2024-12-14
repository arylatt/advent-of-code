package aoc202414

import (
	"math"
	"regexp"
	"slices"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

type Robot struct {
	Start    elves.Point
	Position elves.Point
	Velocity elves.Point
}

const ExprInput = `p=(\d+),(\d+) v=(-?)(\d+),(-?)(\d+)`

var exprInput = regexp.MustCompile(ExprInput)

func Part1(input string) (output string) {
	after100Secs := func(_ []Robot, sec int) bool {
		if sec == 99 {
			return true
		}

		return false
	}

	robots, maxX, maxY := buildGridWithRobots(input)
	robots, _ = move(robots, maxX, maxY, after100Secs)

	quadrants := [][2]elves.Point{
		{elves.Point{X: 0, Y: 0}, elves.Point{X: 49, Y: 50}},
		{elves.Point{X: 51, Y: 0}, elves.Point{X: 100, Y: 50}},
		{elves.Point{X: 0, Y: 52}, elves.Point{X: 49, Y: 102}},
		{elves.Point{X: 51, Y: 52}, elves.Point{X: 100, Y: 102}},
	}

	if elves.IsSample() {
		quadrants = [][2]elves.Point{
			{elves.Point{X: 0, Y: 0}, elves.Point{X: 4, Y: 2}},
			{elves.Point{X: 6, Y: 0}, elves.Point{X: 10, Y: 2}},
			{elves.Point{X: 0, Y: 4}, elves.Point{X: 4, Y: 6}},
			{elves.Point{X: 6, Y: 4}, elves.Point{X: 10, Y: 6}},
		}
	}

	answer := 0

	for _, quad := range quadrants {
		quadAnswer := 0
		for i := 0; i < len(robots); i++ {
			if robots[i].Position.InBounds(quad[0], quad[1]) {
				quadAnswer++
				robots = slices.Delete(robots, i, i+1)
				i--
			}
		}

		if answer == 0 {
			answer = quadAnswer
		} else {
			answer *= quadAnswer
		}
	}

	return strconv.Itoa(answer)
}

func buildGridWithRobots(input string) (robots []Robot, maxX, maxY int) {
	maxX, maxY = 100, 102
	if elves.IsSample() {
		maxX, maxY = 10, 6
	}

	for _, line := range elves.SplitIntoLines(input) {
		position := exprInput.FindStringSubmatch(line)

		x, y := elves.Atoi(position[1]), elves.Atoi(position[2])
		vx, vy := elves.Atoi(position[4]), elves.Atoi(position[6])

		if position[3] == "-" {
			vx = -vx
		}

		if position[5] == "-" {
			vy = -vy
		}

		robots = append(robots, Robot{
			Start:    elves.Point{X: x, Y: y},
			Position: elves.Point{X: x, Y: y},
			Velocity: elves.Point{X: vx, Y: vy},
		})
	}

	return
}

func move(robots []Robot, maxX, maxY int, f func([]Robot, int) bool) ([]Robot, int) {
	sec := 0
	for {
		for i := range robots {
			newPos := robots[i].Position.ShiftPos(robots[i].Velocity)
			if !newPos.Valid(maxX, maxY) {
				if newPos.X < 0 {
					newPos.X = newPos.X + (maxX + 1)
				}

				if newPos.Y < 0 {
					newPos.Y = newPos.Y + (maxY + 1)
				}

				if newPos.X > maxX {
					newPos.X = newPos.X - (maxX + 1)
				}

				if newPos.Y > maxY {
					newPos.Y = newPos.Y - (maxY + 1)
				}
			}

			robots[i].Position = newPos
		}

		if f != nil && f(robots, sec) {
			return robots, sec
		}

		sec++
	}
}

func Part2(input string) (output string) {
	findNearbyRobots := func(robots []Robot, _ int) bool {
		nearbyRobots := 0
		robotPairs := 0
		for i := range robots {
			r1 := robots[i].Position
			for j := range robots {
				if i == j {
					continue
				}

				robotPairs++

				r2 := robots[j].Position

				distance := math.Sqrt(math.Pow(float64(r2.X-r1.X), 2) + math.Pow(float64(r2.Y-r1.Y), 2))

				if distance < 2 {
					nearbyRobots++
				}
			}
		}

		if float64(nearbyRobots/robotPairs)*100 >= 75 {
			return true
		}

		return false
	}

	robots, maxX, maxY := buildGridWithRobots(input)
	_, sec := move(robots, maxX, maxY, findNearbyRobots)
	return strconv.Itoa(sec)
}
