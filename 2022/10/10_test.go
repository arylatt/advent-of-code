package aoc202210

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample1Answer is the expected answer for sample 1
	Sample1Answer = "13140"

	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
)

func TestPart1Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample1Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2022", "10")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample2Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	input, err := elves.GetAOCInput("2022", "10")

	if assert.NoError(t, err) {
		result := Part2(input)

		t.Logf("Result:\r\n%s\r\n", result)
	}
}
