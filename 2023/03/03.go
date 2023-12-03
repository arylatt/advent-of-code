package aoc202303

import (
	"regexp"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

const (
	ExprNum    = `\d+`
	ExprSymbol = `[^.\d\n]`
)

var (
	exprNum    = regexp.MustCompile(ExprNum)
	exprSymbol = regexp.MustCompile(ExprSymbol)
	maxX, maxY int
)

type num struct {
	val   int
	start int
	end   int
	row   int
}

func (n num) symbolPoints() map[elves.Point]bool {
	points := map[elves.Point]bool{}

	if n.start != 0 {
		points[elves.Point{X: n.start - 1, Y: n.row}] = true
	}

	if n.end != maxX {
		points[elves.Point{X: n.end + 1, Y: n.row}] = true
	}

	start := n.start
	end := n.end

	if start != 0 {
		start--
	}

	if end != maxX {
		end++
	}

	if n.row != 0 {
		for i := start; i <= end; i++ {
			points[elves.Point{X: i, Y: n.row - 1}] = true
		}
	}

	if n.row != maxY {
		for i := start; i <= end; i++ {
			points[elves.Point{X: i, Y: n.row + 1}] = true
		}
	}

	return points
}

func parseLine(input string, row int, symbols map[elves.Point]bool) ([]num, map[elves.Point]bool) {
	nums := []num{}

	numMatches := exprNum.FindAllStringSubmatchIndex(input, -1)
	symbolMatches := exprSymbol.FindAllStringSubmatchIndex(input, -1)

	for _, n := range numMatches {
		strNum, _ := strconv.Atoi(input[n[0]:n[1]])
		nums = append(nums, num{
			val:   strNum,
			start: n[0],
			end:   n[1] - 1,
			row:   row,
		})
	}

	for _, s := range symbolMatches {
		symbols[elves.Point{X: s[0], Y: row}] = true
	}

	return nums, symbols
}

func (n num) adjacent(symbols map[elves.Point]bool) bool {
	adjacentSymbols := n.symbolPoints()

	for point := range adjacentSymbols {
		if v, ok := symbols[point]; v && ok {
			return true
		}
	}

	return false
}

func (n num) adjacentStars(symbols map[elves.Point]bool, lines []string, points map[elves.Point][]num) map[elves.Point][]num {
	adjacentSymbols := n.symbolPoints()

	for point := range adjacentSymbols {
		if v, ok := symbols[point]; v && ok {
			if lines[point.Y][point.X] == '*' {
				if _, ok := points[point]; !ok {
					points[point] = []num{}
				}

				points[point] = append(points[point], n)
			}
		}
	}

	return points
}

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)
	maxY = len(lines)
	maxX = len(lines[0])

	numbers := []num{}
	symbols := map[elves.Point]bool{}

	for row, line := range lines {
		n, s := parseLine(line, row, symbols)
		symbols = s

		numbers = append(numbers, n...)
	}

	count := 0

	for _, n := range numbers {
		if n.adjacent(symbols) {
			count += n.val
		}
	}

	return strconv.Itoa(count)
}

func Part2(input string) (output string) {
	lines := elves.SplitIntoLines(input)
	maxY = len(lines)
	maxX = len(lines[0])

	numbers := []num{}
	symbols := map[elves.Point]bool{}

	for row, line := range lines {
		n, s := parseLine(line, row, symbols)
		symbols = s

		numbers = append(numbers, n...)
	}

	points := map[elves.Point][]num{}

	for _, n := range numbers {
		points = n.adjacentStars(symbols, lines, points)
	}

	count := 0

	for _, nums := range points {
		if len(nums) != 2 {
			continue
		}

		count += (nums[0].val * nums[1].val)
	}

	return strconv.Itoa(count)
}
