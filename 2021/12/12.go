package aoc202112

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/2021/sub"
)

func Part1(input string) (output string) {
	answer := 0
	inputs := strings.Split(input, "\r\n")

	caves := &sub.CaveSystem{}
	startCave := &sub.Cave{}
	for _, connection := range inputs {
		caveVals := strings.Split(connection, "-")
		c1, c2 := caves.NewCave(caveVals[0]), caves.NewCave(caveVals[1])
		c1.Connect(c2)
		if c1.ID == "start" {
			startCave = c1
		}
		if c2.ID == "start" {
			startCave = c2
		}
	}

	answer = startCave.Navigate()
	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	answer := 0
	inputs := strings.Split(input, "\r\n")

	sub.CavesPart2 = true
	caves := &sub.CaveSystem{}
	startCave := &sub.Cave{}
	for _, connection := range inputs {
		caveVals := strings.Split(connection, "-")
		c1, c2 := caves.NewCave(caveVals[0]), caves.NewCave(caveVals[1])
		c1.Connect(c2)
		if c1.ID == "start" {
			startCave = c1
		}
		if c2.ID == "start" {
			startCave = c2
		}
	}

	answer = startCave.Navigate()
	return strconv.Itoa(answer)
}
