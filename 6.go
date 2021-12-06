package main

import (
	"strconv"
	"strings"
)

var LanternfishSchool []*Lanternfish = []*Lanternfish{}

type Lanternfish struct {
	Cycle int
}

func (l *Lanternfish) Day() {
	l.Cycle--

	if l.Cycle == -1 {
		l.Cycle = 6
		LanternfishSchool = append(LanternfishSchool, &Lanternfish{Cycle: 8})
	}
}

func Day6Exec(path string, days int) (answer int) {
	input := ParseInputFile(path)[0]
	states := strings.Split(input, ",")

	for _, state := range states {
		cycle, _ := strconv.Atoi(state)
		LanternfishSchool = append(LanternfishSchool, &Lanternfish{Cycle: cycle})
	}

	for i := 1; i <= days; i++ {
		currentLanternfish := make([]*Lanternfish, len(LanternfishSchool))
		copy(currentLanternfish, LanternfishSchool)
		for _, lanternfish := range currentLanternfish {
			lanternfish.Day()
		}
	}

	answer = len(LanternfishSchool)
	return
}
