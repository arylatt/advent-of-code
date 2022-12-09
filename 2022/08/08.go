package aoc202208

import (
	"strconv"
	"strings"
)

type Trees struct {
	Trees [][]int
	UpX   int
	UpY   int
}

func (t Trees) Visible(tree, x, y int) bool {
	if x == 0 || y == 0 {
		return true
	}

	if x == t.UpX || y == t.UpY {
		return true
	}

	visibleN, visibleS, visibleE, visibleW := true, true, true, true

	// Check North
	for i := y - 1; i >= 0; i-- {
		if t.Trees[i][x] >= tree {
			visibleN = false
			break
		}
	}

	// Check South
	for i := y + 1; i <= t.UpY; i++ {
		if t.Trees[i][x] >= tree {
			visibleS = false
			break
		}
	}

	// Check West
	for i := x - 1; i >= 0; i-- {
		if t.Trees[y][i] >= tree {
			visibleW = false
			break
		}
	}

	// Check East
	for i := x + 1; i <= t.UpX; i++ {
		if t.Trees[y][i] >= tree {
			visibleE = false
			break
		}
	}

	return visibleN || visibleS || visibleE || visibleW
}

func (t Trees) Scenic(tree, x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	if x == t.UpX || y == t.UpY {
		return 0
	}

	sN, sS, sE, sW := 0, 0, 0, 0

	// Check North
	for i := y - 1; i >= 0; i-- {
		sN++
		if t.Trees[i][x] >= tree {
			break
		}
	}

	// Check South
	for i := y + 1; i <= t.UpY; i++ {
		sS++
		if t.Trees[i][x] >= tree {
			break
		}
	}

	// Check West
	for i := x - 1; i >= 0; i-- {
		sW++
		if t.Trees[y][i] >= tree {
			break
		}
	}

	// Check East
	for i := x + 1; i <= t.UpX; i++ {
		sE++
		if t.Trees[y][i] >= tree {
			break
		}
	}

	return sN * sS * sE * sW
}

func ParseTrees(input []string) (trees Trees) {
	trees = Trees{Trees: [][]int{}}

	for _, row := range input {
		rowTrees := []int{}

		for _, tree := range strings.Split(row, "") {
			treeInt, _ := strconv.Atoi(tree)
			rowTrees = append(rowTrees, treeInt)
		}

		trees.Trees = append(trees.Trees, rowTrees)
	}

	trees.UpX = len(trees.Trees[0]) - 1
	trees.UpY = len(trees.Trees) - 1

	return
}

func Part1(input string) (output string) {
	trees := ParseTrees(strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n"))
	total := 0

	for y, row := range trees.Trees {
		for x, tree := range row {
			if trees.Visible(tree, x, y) {
				total++
			}
		}
	}

	return strconv.Itoa(total)
}

func Part2(input string) (output string) {
	trees := ParseTrees(strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n"))
	max := 0

	for y, row := range trees.Trees {
		for x, tree := range row {
			if scenic := trees.Scenic(tree, x, y); scenic > max {
				max = scenic
			}
		}
	}

	return strconv.Itoa(max)
}
