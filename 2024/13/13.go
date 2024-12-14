package aoc202413

import (
	"regexp"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

const (
	ExprA           = `Button A: X\+(\d+), Y\+(\d+)`
	ExprB           = `Button B: X\+(\d+), Y\+(\d+)`
	ExprPrize       = `Prize: X=(\d+), Y=(\d+)`
	CoordCorrection = 10000000000000
)

var (
	exprA     = regexp.MustCompile(ExprA)
	exprB     = regexp.MustCompile(ExprB)
	exprPrize = regexp.MustCompile(ExprPrize)
)

func Part1(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	answer := 0

	for i := 0; i < len(lines); i += 3 {
		a := exprA.FindStringSubmatch(lines[i])
		b := exprB.FindStringSubmatch(lines[i+1])
		prize := exprPrize.FindStringSubmatch(lines[i+2])

		aX, aY := elves.Atoi(a[1]), elves.Atoi(a[2])
		bX, bY := elves.Atoi(b[1]), elves.Atoi(b[2])
		pX, pY := elves.Atoi(prize[1]), elves.Atoi(prize[2])

		// equations will be
		// (aX.A) + (bX.B) = pX
		// (aY.A) + (bY.B) = pY

		aX1, pX1 := aX*bY, pX*bY
		aY1, pY1 := aY*bX, pY*bX

		aTotal := aX1 - aY1
		pTotal := pX1 - pY1

		aVal := pTotal / aTotal
		bVal := (pX - (aX * aVal)) / bX

		if (aX*aVal)+(bX*bVal) == pX && (aY*aVal)+(bY*bVal) == pY {
			answer += (aVal * 3) + bVal
		}
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	lines := elves.SplitIntoLines(input)

	answer := 0

	for i := 0; i < len(lines); i += 3 {
		a := exprA.FindStringSubmatch(lines[i])
		b := exprB.FindStringSubmatch(lines[i+1])
		prize := exprPrize.FindStringSubmatch(lines[i+2])

		aX, aY := elves.Atoi(a[1]), elves.Atoi(a[2])
		bX, bY := elves.Atoi(b[1]), elves.Atoi(b[2])
		pX, pY := elves.Atoi(prize[1])+CoordCorrection, elves.Atoi(prize[2])+CoordCorrection

		// equations will be
		// (aX.A) + (bX.B) = pX
		// (aY.A) + (bY.B) = pY

		aX1, pX1 := aX*bY, pX*bY
		aY1, pY1 := aY*bX, pY*bX

		aTotal := aX1 - aY1
		pTotal := pX1 - pY1

		aVal := pTotal / aTotal
		bVal := (pX - (aX * aVal)) / bX

		if (aX*aVal)+(bX*bVal) == pX && (aY*aVal)+(bY*bVal) == pY {
			answer += (aVal * 3) + bVal
		}
	}

	return strconv.Itoa(answer)
}
