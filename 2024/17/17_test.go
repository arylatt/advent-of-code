package aoc202417

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample1Answer is the expected answer for sample 1
	Sample1Answer = "4,6,3,5,6,3,5,2,1,0"

	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "117440"
)

func TestPart1Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample1Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2024", "17")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample2Answer, "testdata/sample2.txt")

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2024", "17")
}