package sub

import (
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func ParsePosition(coord string) (pos Position) {
	data := strings.Split(coord, ",")
	pos.X, _ = strconv.Atoi(data[0])
	pos.Y, _ = strconv.Atoi(data[1])
	return
}
