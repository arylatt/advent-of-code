package aoc202207

import (
	"strings"
	"testing"

	"github.com/arylatt/advent-of-code/elves"
	"github.com/stretchr/testify/assert"
)

const (
	// Sample1Answer is the expected answer for sample 1
	Sample1Answer = "95437"

	// Sample2Answer is the expected answer for sample 2
	Sample2Answer = "24933642"
)

func TestPart1Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample1Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part1)
	}
}

func TestPart1(t *testing.T) {
	elves.TestReal(t, Part1, 1, "2022", "7")
}

func TestPart2Sample(t *testing.T) {
	td, err := elves.SampleFileToTestData(Sample2Answer)

	if assert.NoError(t, err) {
		elves.TestSample(t, td, Part2)
	}
}

func TestPart2(t *testing.T) {
	elves.TestReal(t, Part2, 2, "2022", "7")
}

func TestWalk(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

	Walk(strings.Split(input, "\n"))
}
