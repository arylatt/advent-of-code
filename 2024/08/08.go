package aoc202408

import (
	"math"
	"slices"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

func Part1(input string) (output string) {
	antennas, maxX, maxY := parseMap(input)

	antinodes := []elves.Point{}

	for _, points := range antennas {
		antinodes = append(antinodes, findAntinodes(points, maxX, maxY, false)...)
	}

	antinodesMap := map[elves.Point]int{}
	for _, antinode := range antinodes {
		antinodesMap[antinode] = 0
	}

	return strconv.Itoa(len(antinodesMap))
}

func parseMap(input string) (map[string][]elves.Point, int, int) {
	lines := elves.SplitIntoLines(input)

	maxX, maxY := len(lines[0])-1, len(lines)-1
	antennas := map[string][]elves.Point{}

	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}

			if _, ok := antennas[string(char)]; !ok {
				antennas[string(char)] = []elves.Point{}
			}

			antennas[string(char)] = append(antennas[string(char)], elves.Point{X: x, Y: y})
		}
	}

	return antennas, maxX, maxY
}

func findAntinodes(antennas []elves.Point, maxX, maxY int, part2 bool) []elves.Point {
	antinodes := []elves.Point{}

	for _, antennaA := range antennas {
		if !slices.Contains(antinodes, antennaA) && part2 {
			antinodes = append(antinodes, antennaA)
		}

		for _, antennaB := range antennas {
			if antennaA.Equals(antennaB) {
				continue
			}

			antinodeX, antinodeY := int(math.Abs(float64(antennaA.X-antennaB.X))), int(math.Abs(float64(antennaA.Y-antennaB.Y)))

			if antennaB.X < antennaA.X {
				antinodeX = antennaB.X - antennaA.X
			}

			if antennaB.Y < antennaA.Y {
				antinodeY = antennaB.Y - antennaA.Y
			}

			antinode := antennaB

			for {
				antinode = antinode.Shift(antinodeX, antinodeY)
				if !antinode.Valid(maxX, maxY) {
					break
				}

				if !slices.Contains(antinodes, antinode) {
					antinodes = append(antinodes, antinode)
				}

				if !part2 {
					break
				}
			}
		}
	}

	return antinodes
}

func Part2(input string) (output string) {
	antennas, maxX, maxY := parseMap(input)

	antinodes := []elves.Point{}

	for _, points := range antennas {
		antinodes = append(antinodes, findAntinodes(points, maxX, maxY, true)...)
	}

	antinodesMap := map[elves.Point]int{}
	for _, antinode := range antinodes {
		antinodesMap[antinode] = 0
	}

	return strconv.Itoa(len(antinodesMap))
}
