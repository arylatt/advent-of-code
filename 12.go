package main

import (
	"strings"

	"github.com/arylatt/advent-of-code/sub"
)

func Day12Exec(path string) (answer int) {
	inputs := ParseInputFile(path)

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

	return
}

func Day12ExecII(path string) (answer int) {
	inputs := ParseInputFile(path)

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

	return
}
