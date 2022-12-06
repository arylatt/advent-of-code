package sub

import (
	"strconv"
	"strings"
)

type Command struct {
	Direction Direction
	Units     int
}

func (c *Command) Parse(input string) {
	cmdData := strings.Split(input, " ")

	switch cmdData[0] {
	case "forward":
		c.Direction = Forward
	case "up":
		c.Direction = Up
	case "down":
		c.Direction = Down
	}

	c.Units, _ = strconv.Atoi(strings.TrimSpace(cmdData[1]))
}
