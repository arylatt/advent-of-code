package bingo

import (
	"regexp"
	"strconv"
	"strings"
)

type Row [5]int

type mark [5]bool

type Card struct {
	Rows [5]Row

	marks [5]mark
}

func GenerateCard(inputs [5]string) *Card {
	c := &Card{}

	regex := regexp.MustCompile(`\s+`)

	for i, input := range inputs {
		vals := regex.Split(strings.TrimSpace(input), -1)
		for j, val := range vals {
			c.Rows[i][j], _ = strconv.Atoi(val)
		}
	}

	return c
}

func GenerateCards(inputs []string) []*Card {
	input := [5]string{}
	cards := []*Card{}
	i := 0

	for _, str := range inputs {
		if str == "" {
			continue
		}

		input[i] = str
		i++

		if i >= 5 {
			cards = append(cards, GenerateCard(input))
			input = [5]string{}
			i = 0
		}
	}

	return cards
}

func (c *Card) Mark(input int) bool {
	for i, row := range c.Rows {
		for j, val := range row {
			if val == input {
				c.marks[i][j] = true
				return true
			}
		}
	}

	return false
}

func MarkCards(cards []*Card, input int) (marked []*Card) {
	for _, card := range cards {
		mark := card.Mark(input)
		if mark {
			marked = append(marked, card)
		}
	}

	return
}

func (c *Card) HasWonRow(row int) bool {
	for _, mark := range c.marks[row] {
		if !mark {
			return false
		}
	}

	return true
}

func (c *Card) HasWonColumn(column int) bool {
	for row := range c.marks {
		if !c.marks[row][column] {
			return false
		}
	}

	return true
}

func (c *Card) HasWon() bool {
	for i := 0; i < 5; i++ {
		if c.HasWonRow(i) || c.HasWonColumn(i) {
			return true
		}
	}

	return false
}

func GetWinningCards(cards []*Card) (winners []*Card) {
	for _, card := range cards {
		if card.HasWon() {
			winners = append(winners, card)
		}
	}
	return
}

func GetOrderedWinningCards(cards []*Card) (winners []*Card) {
	winners = make([]*Card, len(cards))

	for i, card := range cards {
		if card.HasWon() {
			winners[i] = card
		}
	}
	return
}

func (c *Card) SumOfUnmarkedNumbers() (result int) {
	for row := range c.marks {
		for col, mark := range c.marks[row] {
			if !mark {
				result += c.Rows[row][col]
			}
		}
	}

	return
}
