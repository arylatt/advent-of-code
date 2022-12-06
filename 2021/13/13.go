package aoc202113

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/2021/sub"
)

func Part1(input string) (output string) {
	answer := 0
	inputs := strings.Split(input, "\r\n")

	points := []sub.Position{}
	folds := []sub.Position{}
	parsePoints := true

	for _, input := range inputs {
		if parsePoints {
			if input == "" {
				parsePoints = false
				continue
			}

			points = append(points, sub.ParsePosition(input))
			continue
		}

		input = strings.ReplaceAll(input, "fold along ", "")
		inputParts := strings.Split(input, "=")
		foldPos, _ := strconv.Atoi(inputParts[1])
		switch inputParts[0] {
		case "x":
			folds = append(folds, sub.Position{X: foldPos})
		case "y":
			folds = append(folds, sub.Position{Y: foldPos})
		}
	}

	highX, highY := 0, 0
	for _, point := range points {
		if point.X >= highX {
			highX = point.X + 1
		}
		if point.Y >= highY {
			highY = point.Y + 1
		}
	}

	paper := make([][]int, highY)
	for y := range paper {
		paper[y] = make([]int, highX)
	}

	for _, p := range points {
		paper[p.Y][p.X] = 1
	}

	for _, fold := range folds[0:1] {
		if fold.X == 0 {
			folded := make([][]int, fold.Y)
			copy(folded, paper[:fold.Y])

			for i := 0; i < fold.Y; i++ {
				for x, val := range paper[(len(paper)-1)-i] {
					if val == 1 {
						folded[i][x] = 1
					}
				}
			}

			paper = folded
			continue
		}

		folded := make([][]int, len(paper))
		for y := range folded {
			folded[y] = make([]int, fold.X)
			copy(folded[y], paper[y][:fold.X])

			for i := 0; i < fold.X; i++ {
				if paper[y][(len(paper[y])-1)-i] == 1 {
					folded[y][i] = 1
				}
			}
		}
		paper = folded
	}

	for y := range paper {
		for x := range paper[y] {
			answer += paper[y][x]
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	answer := 0
	inputs := strings.Split(input, "\r\n")

	points := []sub.Position{}
	folds := []sub.Position{}
	parsePoints := true

	for _, input := range inputs {
		if parsePoints {
			if input == "" {
				parsePoints = false
				continue
			}

			points = append(points, sub.ParsePosition(input))
			continue
		}

		input = strings.ReplaceAll(input, "fold along ", "")
		inputParts := strings.Split(input, "=")
		foldPos, _ := strconv.Atoi(inputParts[1])
		switch inputParts[0] {
		case "x":
			folds = append(folds, sub.Position{X: foldPos})
		case "y":
			folds = append(folds, sub.Position{Y: foldPos})
		}
	}

	highX, highY := 0, 0
	for _, point := range points {
		if point.X >= highX {
			highX = point.X + 1
		}
		if point.Y >= highY {
			highY = point.Y + 1
		}
	}

	paper := make([][]int, highY)
	for y := range paper {
		paper[y] = make([]int, highX)
	}

	for _, p := range points {
		paper[p.Y][p.X] = 1
	}

	for _, fold := range folds {
		if fold.X == 0 {
			folded := make([][]int, fold.Y)
			copy(folded, paper[:fold.Y])

			for i := 0; i < fold.Y; i++ {
				for x, val := range paper[(len(paper)-1)-i] {
					if val == 1 {
						folded[i][x] = 1
					}
				}
			}

			paper = folded
			continue
		}

		folded := make([][]int, len(paper))
		for y := range folded {
			folded[y] = make([]int, fold.X)
			copy(folded[y], paper[y][:fold.X])

			for i := 0; i < fold.X; i++ {
				if paper[y][(len(paper[y])-1)-i] == 1 {
					folded[y][i] = 1
				}
			}
		}
		paper = folded
	}

	fmt.Println()
	for y := range paper {
		for x := range paper[y] {
			answer += paper[y][x]
			if paper[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	return
}
