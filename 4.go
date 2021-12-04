package main

import (
	"strconv"
	"strings"

	"github.com/arylatt/advent-of-code/bingo"
)

func Day4Exec(path string) (answer int) {
	inputs := ParseInputFile(path)

	cards := bingo.GenerateCards(inputs[2:])
	numbers := strings.Split(inputs[0], ",")

	for _, val := range numbers {
		num, _ := strconv.Atoi(val)

		bingo.MarkCards(cards, num)

		winningCards := bingo.GetWinningCards(cards)
		if len(winningCards) != 0 {
			answer = num * winningCards[0].SumOfUnmarkedNumbers()
			return
		}
	}

	return
}

func Day4ExecII(path string) (answer int) {
	inputs := ParseInputFile(path)

	cards := bingo.GenerateCards(inputs[2:])
	numbers := strings.Split(inputs[0], ",")

	lastNumberChange := 0
	lastWinningCard := []*bingo.Card{{}}
	winningCardIndexes := []int{}

	for _, val := range numbers {
		num, _ := strconv.Atoi(val)

		bingo.MarkCards(cards, num)

		winningCards := bingo.GetOrderedWinningCards(cards)
		for i, card := range winningCards {
			if card == nil {
				continue
			}

			winnerIndexed := false
			for _, winner := range winningCardIndexes {
				if i == winner {
					winnerIndexed = true
					break
				}
			}

			if winnerIndexed {
				continue
			}

			winningCardIndexes = append(winningCardIndexes, i)
			lastNumberChange = num

			copy(lastWinningCard, []*bingo.Card{card})
		}

		if len(cards) == len(bingo.GetWinningCards(cards)) {
			break
		}
	}

	answer = lastNumberChange * lastWinningCard[0].SumOfUnmarkedNumbers()

	return
}
