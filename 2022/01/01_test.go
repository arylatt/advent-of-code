package aoc202201

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	file, _ := os.ReadFile("testdata/sample.txt")

	assert.Equal(t, "24000", Part1(string(file)))
}

func TestPart1(t *testing.T) {
	file, _ := os.ReadFile("testdata/input.txt")

	assert.Equal(t, "68775", Part1(string(file)))
}

func TestPart2Sample(t *testing.T) {
	file, _ := os.ReadFile("testdata/sample.txt")

	assert.Equal(t, "45000", Part2(string(file)))
}

func TestPart2(t *testing.T) {
	file, _ := os.ReadFile("testdata/input.txt")

	assert.Equal(t, "202585", Part2(string(file)))
}
