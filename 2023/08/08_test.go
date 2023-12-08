package aoc202308

import (
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "6"
)

func TestPart1Sample1(t *testing.T) {
	ans1, ans2 := "2", "6"
	td, err := elves.SampleFileToTestData(ans1)
	td2, err2 := elves.SampleFileToTestData(ans2, "testdata/sample2.txt")

	if assert.NoError(t, err) && assert.NoError(t, err2) {
		elves.TestSample(t, td, Part1)
		elves.TestSample(t, td2, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2023", "8")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample2Answer, "testdata/sample3.txt")

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2023", "8")
}
