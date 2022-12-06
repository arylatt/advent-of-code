package aoc202109

import (
	"sort"
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/2021/sub"
)

func posInBasin(pos *sub.Position, basin []*sub.Position) bool {
	for _, p := range basin {
		if p.X == pos.X && p.Y == pos.Y {
			return true
		}
	}
	return false
}

func Part1(input string) (output string) {
	answer := 0
	inputs := strings.Split(input, "\r\n")
	heightMap := make([][]int, len(inputs))

	for i, input := range inputs {
		heightMap[i] = make([]int, len(input))
		for j, val := range strings.Split(input, "") {
			v, _ := strconv.Atoi(val)
			heightMap[i][j] = v
		}
	}

	rowLen, colLen := len(heightMap)-1, len(heightMap[0])-1

	for i, row := range heightMap {
		for j, val := range row {
			above, left, right, below := val+1, val+1, val+1, val+1
			if i != 0 {
				above = heightMap[i-1][j]
			}
			if j != 0 {
				left = heightMap[i][j-1]
			}
			if i != rowLen {
				below = heightMap[i+1][j]
			}
			if j != colLen {
				right = heightMap[i][j+1]
			}

			if val < above && val < left && val < below && val < right {
				answer += val + 1
			}
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	answer := 0
	inputs := strings.Split(input, "\r\n")
	heightMap := make([][]int, len(inputs))

	for i, input := range inputs {
		heightMap[i] = make([]int, len(input))
		for j, val := range strings.Split(input, "") {
			v, _ := strconv.Atoi(val)
			heightMap[i][j] = v
		}
	}

	basins := [][]*sub.Position{}

	rowLen, colLen := len(heightMap)-1, len(heightMap[0])-1

	for i, row := range heightMap {
		for j, val := range row {
			above, left, right, below := val+1, val+1, val+1, val+1
			if i != 0 {
				above = heightMap[i-1][j]
			}
			if j != 0 {
				left = heightMap[i][j-1]
			}
			if i != rowLen {
				below = heightMap[i+1][j]
			}
			if j != colLen {
				right = heightMap[i][j+1]
			}

			if val < above && val < left && val < below && val < right {
				basins = append(basins, []*sub.Position{{X: j, Y: i}})
			}
		}
	}

	for i, basin := range basins {
		basinSize := 0
		for basinSize != len(basin) {
			basinSize = len(basin)
			for _, b := range basin {
				cVal := heightMap[b.Y][b.X]
				if b.Y != 0 {
					nPos := &sub.Position{Y: b.Y - 1, X: b.X}
					if !posInBasin(nPos, basin) {
						nVal := heightMap[nPos.Y][nPos.X]
						if nVal != 9 && nVal > cVal {
							basin = append(basin, nPos)
						}
					}
				}
				if b.X != 0 {
					nPos := &sub.Position{Y: b.Y, X: b.X - 1}
					if !posInBasin(nPos, basin) {
						nVal := heightMap[nPos.Y][nPos.X]
						if nVal != 9 && nVal > cVal {
							basin = append(basin, nPos)
						}
					}
				}
				if b.Y != rowLen {
					nPos := &sub.Position{Y: b.Y + 1, X: b.X}
					if !posInBasin(nPos, basin) {
						nVal := heightMap[nPos.Y][nPos.X]
						if nVal != 9 && nVal > cVal {
							basin = append(basin, nPos)
						}
					}
				}
				if b.X != colLen {
					nPos := &sub.Position{Y: b.Y, X: b.X + 1}
					if !posInBasin(nPos, basin) {
						nVal := heightMap[nPos.Y][nPos.X]
						if nVal != 9 && nVal > cVal {
							basin = append(basin, nPos)
						}
					}
				}
			}
		}
		basins[i] = basin
	}

	answers := []int{}

	for _, b := range basins {
		answers = append(answers, len(b))
	}

	sort.Ints(answers)
	answersLen := len(answers) - 1

	answer = answers[answersLen] * answers[answersLen-1] * answers[answersLen-2]

	return strconv.Itoa(answer)
}
