package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Lanternfish struct {
	Cycle int
}

type LanternfishResult struct {
	StartCycle int
	StartDays  int
	Result     int64
}

var resultsMutex *sync.Mutex = &sync.Mutex{}
var results []*LanternfishResult = []*LanternfishResult{}

func GetLanternfishResult(cycle, days int) (result *int64) {
	resultsMutex.Lock()
	defer resultsMutex.Unlock()

	for _, res := range results {
		if res.StartCycle == cycle && res.StartDays == days {
			result = &res.Result
			break
		}
	}

	return
}

func AddLanternfishResult(cycle, days int, result int64) {
	resultsMutex.Lock()
	defer resultsMutex.Unlock()

	for _, res := range results {
		if res.StartCycle == cycle && res.StartDays == days {
			if res.Result != result {
				panic(fmt.Sprintf("duplicate result %v with different answer to %d", res, result))
			}
		}
	}

	results = append(results, &LanternfishResult{StartCycle: cycle, StartDays: days, Result: result})
}

func (l *Lanternfish) Run(results chan int64, days int) {
	result := GetLanternfishResult(l.Cycle, days)
	if result != nil {
		results <- *result
		return
	}

	startCycle := l.Cycle
	res := int64(1)

	for i := 1; i <= days; i++ {
		l.Cycle--

		if l.Cycle == -1 {
			l.Cycle = 6
			await := make(chan int64, 1)
			go (&Lanternfish{Cycle: 8}).Run(await, days-i)
			res += <-await
		}
	}

	AddLanternfishResult(startCycle, days, res)

	results <- res
}

func Day6Exec(path string, days int) (answer int64) {
	input := ParseInputFile(path)[0]
	states := strings.Split(input, ",")

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
		results := make(chan int64, 1)
		go (&Lanternfish{Cycle: state}).Run(results, days)
		answers[state-1] = <-results
	}

	for _, state := range states {
		stateInt, _ := strconv.Atoi(state)
		answer += answers[stateInt-1]
	}

	return
}
