package aoc202211

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	ID          int
	Items       []uint64
	Operation   func(old uint64) (new uint64)
	TestDiv     uint64
	TrueMonkey  int
	FalseMonkey int
	Inspected   uint64
}

func NewMonkey(input []string) (m *Monkey) {
	m = &Monkey{}

	m.ID, _ = strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(input[0]), "Monkey "), ":"))

	startItemsStr := strings.Split(strings.TrimPrefix(strings.TrimSpace(input[1]), "Starting items: "), ", ")
	for _, startItemStr := range startItemsStr {
		startItem, _ := strconv.ParseUint(startItemStr, 10, 64)
		m.Items = append(m.Items, startItem)
	}

	opsStr := strings.Split(strings.TrimPrefix(strings.TrimSpace(input[2]), "Operation: new = "), " ")
	m.Operation = func(old uint64) (new uint64) {
		op1, _ := strconv.ParseUint(opsStr[0], 10, 64)
		if opsStr[0] == "old" {
			op1 = old
		}

		op2, _ := strconv.ParseUint(opsStr[2], 10, 64)
		if opsStr[2] == "old" {
			op2 = old
		}

		switch opsStr[1] {
		case "*":
			return op1 * op2
		case "+":
			return op1 + op2
		case "/":
			return op1 / op2
		case "-":
			return op1 - op2
		}

		return 0
	}

	m.TestDiv, _ = strconv.ParseUint(strings.TrimPrefix(strings.TrimSpace(input[3]), "Test: divisible by "), 10, 64)
	m.TrueMonkey, _ = strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(input[4]), "If true: throw to monkey "))
	m.FalseMonkey, _ = strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(input[5]), "If false: throw to monkey "))

	return
}

func (m *Monkey) Inspect() {
	for len(m.Items) != 0 {
		m.Inspected++

		item := m.Operation(m.Items[0])

		if worryManaged {
			item = uint64(math.Floor(float64(item) / 3))
		} else {
			// TODO: hmm...
		}

		if item%m.TestDiv == 0 {
			monkeys[m.TrueMonkey].Items = append(monkeys[m.TrueMonkey].Items, item)
		} else {
			monkeys[m.FalseMonkey].Items = append(monkeys[m.FalseMonkey].Items, item)
		}

		m.Items = m.Items[1:]
	}
}

var monkeys = []*Monkey{}
var worryManaged = true

func Chase(input string, rounds int) {
	inputs := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")

	for i := 0; i < len(inputs); i += 7 {
		monkeys = append(monkeys, NewMonkey(inputs[i:i+6]))
	}

	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			m.Inspect()
		}
	}
}

func Part1(input string) (output string) {
	worryManaged = true

	Chase(input, 20)

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})

	return strconv.FormatUint(monkeys[0].Inspected*monkeys[1].Inspected, 10)
}

func Part2(input string) (output string) {
	worryManaged = false

	Chase(input, 10000)

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})

	return strconv.FormatUint(monkeys[0].Inspected*monkeys[1].Inspected, 10)
}
