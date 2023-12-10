package aoc202309

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample1Answer is the expected answer for sample 1
	Sample1Answer = "114"

	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "2"
)

func TestExtrapolateForward(t *testing.T) {
	expected := [][]int{
		{0, 3, 6, 9, 12, 15, 18},
		{3, 3, 3, 3, 3, 3},
		{0, 0, 0, 0},
	}

	assert.Equal(t, expected, extrapolateForward("0 3 6 9 12 15"))
}

func TestPart1Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample1Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2023", "9")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample2Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2023", "9")
}
