package aoc202106

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Lanternfish struct {
	Cycle int
}

type LanternfishResult struct {
	StartCycle int
	StartDays  int
	Result     int64
}

var results []*LanternfishResult = []*LanternfishResult{}

func GetLanternfishResult(cycle, days int) (result *int64) {
	for _, res := range results {
		if res.StartCycle == cycle && res.StartDays == days {
			result = &res.Result
			break
		}
	}

	return
}

func AddLanternfishResult(cycle, days int, result int64) {
	for _, res := range results {
		if res.StartCycle == cycle && res.StartDays == days {
			if res.Result != result {
				panic(fmt.Sprintf("duplicate result %v with different answer to %d", res, result))
			}
		}
	}

	results = append(results, &LanternfishResult{StartCycle: cycle, StartDays: days, Result: result})
}

func (l *Lanternfish) Run(days int) (result int64) {
	solved := GetLanternfishResult(l.Cycle, days)
	if solved != nil {
		result = *solved
		return
	}

	startCycle := l.Cycle
	result = 1

	for i := 1; i <= days; i++ {
		l.Cycle--

		if l.Cycle == -1 {
			l.Cycle = 6
			result += (&Lanternfish{Cycle: 8}).Run(days - i)
		}
	}

	AddLanternfishResult(startCycle, days, result)
	return
}

func Execute(input string, days int) (answer int64) {
	inputs := strings.Split(input, "\r\n")[0]
	states := strings.Split(inputs, ",")

	uniqueStates := []int{}
	for i, state := range states {
		found := false
		for _, uniqueState := range states[:i] {
			if state == uniqueState {
				found = true
				break
			}
		}

		if !found {
			stateInt, _ := strconv.Atoi(state)
			uniqueStates = append(uniqueStates, stateInt)
		}
	}

	sort.Ints(uniqueStates)

	answers := make([]int64, uniqueStates[len(uniqueStates)-1])
	for _, state := range uniqueStates {
		answers[state-1] = (&Lanternfish{Cycle: state}).Run(days)
	}

	for _, state := range states {
		stateInt, _ := strconv.Atoi(state)
		answer += answers[stateInt-1]
	}

	return
}

func Part1(input string) (output string) {
	return strconv.Itoa(int(Execute(input, 80)))
}

func Part2(input string) (output string) {
	return strconv.FormatInt((Execute(input, 256)), 10)
}
