package sevenseg

import "strings"

const (
	Top = iota
	TopLeft
	TopRight
	Middle
	BottomLeft
	BottomRight
	Bottom
)

type Display [7]string

func GeneratePossibleDisplays(input string, inputDisplays []*Display) (displays []*Display) {
	switch len(input) {
	case 2:
		d1 := &Display{TopLeft: string(input[0]), BottomLeft: string(input[1])}
		d2 := &Display{TopLeft: string(input[1]), BottomLeft: string(input[0])}
		displays = append(displays, d1, d2)
		return
	case 3:
		for _, displayChar := range inputDisplays[0] {
			input = strings.ReplaceAll(input, displayChar, "")
		}

		for _, display := range inputDisplays {
			display[Top] = input
			displays = append(displays, display)
		}

		return
	}

	return
}
