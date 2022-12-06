package sub

import (
	"os"
	"strings"
)

const (
	SmallCave = iota
	BigCave
	StartCave
	EndCave
)

type Cave struct {
	ID          string
	Connections []*Cave
	Type        int
}

type CaveSystem struct {
	Caves []*Cave
}

var CavesPart2 bool = false

func (cs *CaveSystem) NewCave(id string) (c *Cave) {
	c = &Cave{ID: id}

	switch strings.ToLower(c.ID) {
	case "start":
		c.Type = StartCave
	case "end":
		c.Type = EndCave
	case id:
		c.Type = SmallCave
	default:
		c.Type = BigCave
	}

	for _, ec := range cs.Caves {
		if ec.ID == c.ID && ec.Type == c.Type {
			return ec
		}
	}

	cs.Caves = append(cs.Caves, c)

	return
}

func (c *Cave) Connect(c1 *Cave) {
	for _, cave := range c.Connections {
		if c1.ID == cave.ID {
			return
		}
	}

	c.Connections = append(c.Connections, c1)
	c1.Connections = append(c1.Connections, c)
}

func (c *Cave) Navigate(history ...string) (answer int) {
	if len(history) == 0 && c.Type != StartCave {
		panic("not start cave")
	}

	if c.Type == EndCave {
		f, _ := os.ReadFile("inputs/12_expect.txt")
		histJoin := strings.Join(append(history, c.ID), ",")
		rm := -1
		lines := strings.Split(string(f), "\r\n")
		for i, l := range lines {
			if l == histJoin {
				rm = i
			}
		}
		if rm > 0 {
			lines = append(lines[:rm], lines[rm+1:]...)
			os.WriteFile("inputs/12_expect.txt", []byte(strings.Join(lines, "\r\n")), 0644)
		}
		return 1
	}

	history = append(history, c.ID)

	for _, conn := range c.Connections {
		if conn.CanVisit(history...) {
			answer += conn.Navigate(history...)
		}
	}

	return
}

func (c *Cave) CanVisit(history ...string) bool {
	switch c.Type {
	case StartCave:
		return false
	case BigCave:
		return true
	case EndCave:
		return true
	case SmallCave:
		return c.canVisitSmallCave(history...)
	}
	return false
}

func (c *Cave) canVisitSmallCave(history ...string) bool {
	if !CavesPart2 {
		for _, visit := range history {
			if visit == c.ID {
				return false
			}
		}
		return true
	}

	// Allow visiting the first small cave twice?
	visits := map[string]int{}
	for _, visit := range history {
		if strings.ToLower(visit) == visit {
			visits[visit]++
		}
	}

	for v := range visits {
		if visits[v] == 2 {
			if visits[c.ID] >= 1 {
				return false
			}
		}
	}

	return true
}
