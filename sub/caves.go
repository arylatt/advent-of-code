package sub

import "strings"

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

func NewCave(id string) (c *Cave) {
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

func (c *Cave) Navigate() (answer int) {
	return
}
