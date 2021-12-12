package main

import (
	"strconv"
	"strings"
)

func Day11Exec(path string) (answer int) {
	inputs := ParseInputFile(path)
	octopi := make([][]int, len(inputs))

	for y, input := range inputs {
		octopi[y] = make([]int, len(input))
		for x, val := range strings.Split(input, "") {
			octopi[y][x], _ = strconv.Atoi(val)
		}
	}

	maxY := len(octopi) - 1
	maxX := len(octopi[0]) - 1

	for i := 0; i < 100; i++ {
		for y := range octopi {
			for x := range octopi[y] {
				octopi[y][x]++
			}
		}

		changed := -1
		for changed != 0 {
			changed = 0
			for y := range octopi {
				for x := range octopi[y] {
					if octopi[y][x] == -1 {
						continue
					}

					if octopi[y][x] > 9 {
						changed++
						octopi[y][x] = -1
						if x != 0 {
							if octopi[y][x-1] != -1 {
								octopi[y][x-1]++
							}
						}
						if x != maxX {
							if octopi[y][x+1] != -1 {
								octopi[y][x+1]++
							}
						}
						if y != 0 {
							if octopi[y-1][x] != -1 {
								octopi[y-1][x]++
							}
						}
						if y != maxY {
							if octopi[y+1][x] != -1 {
								octopi[y+1][x]++
							}
						}
						if x != 0 && y != 0 {
							if octopi[y-1][x-1] != -1 {
								octopi[y-1][x-1]++
							}
						}
						if x != maxX && y != maxY {
							if octopi[y+1][x+1] != -1 {
								octopi[y+1][x+1]++
							}
						}
						if x != 0 && y != maxY {
							if octopi[y+1][x-1] != -1 {
								octopi[y+1][x-1]++
							}
						}
						if x != maxX && y != 0 {
							if octopi[y-1][x+1] != -1 {
								octopi[y-1][x+1]++
							}
						}
					}
				}
			}
		}

		for y := range octopi {
			for x := range octopi[y] {
				if octopi[y][x] == -1 {
					answer++
					octopi[y][x] = 0
				}
			}
		}
	}

	return
}

func Day11ExecII(path string) (answer int) {
	inputs := ParseInputFile(path)
	octopi := make([][]int, len(inputs))

	for y, input := range inputs {
		octopi[y] = make([]int, len(input))
		for x, val := range strings.Split(input, "") {
			octopi[y][x], _ = strconv.Atoi(val)
		}
	}

	maxY := len(octopi) - 1
	maxX := len(octopi[0]) - 1

	for i := 0; answer == 0; i++ {
		for y := range octopi {
			for x := range octopi[y] {
				octopi[y][x]++
			}
		}

		changed := -1
		for changed != 0 {
			changed = 0
			for y := range octopi {
				for x := range octopi[y] {
					if octopi[y][x] == -1 {
						continue
					}

					if octopi[y][x] > 9 {
						changed++
						octopi[y][x] = -1
						if x != 0 {
							if octopi[y][x-1] != -1 {
								octopi[y][x-1]++
							}
						}
						if x != maxX {
							if octopi[y][x+1] != -1 {
								octopi[y][x+1]++
							}
						}
						if y != 0 {
							if octopi[y-1][x] != -1 {
								octopi[y-1][x]++
							}
						}
						if y != maxY {
							if octopi[y+1][x] != -1 {
								octopi[y+1][x]++
							}
						}
						if x != 0 && y != 0 {
							if octopi[y-1][x-1] != -1 {
								octopi[y-1][x-1]++
							}
						}
						if x != maxX && y != maxY {
							if octopi[y+1][x+1] != -1 {
								octopi[y+1][x+1]++
							}
						}
						if x != 0 && y != maxY {
							if octopi[y+1][x-1] != -1 {
								octopi[y+1][x-1]++
							}
						}
						if x != maxX && y != 0 {
							if octopi[y-1][x+1] != -1 {
								octopi[y-1][x+1]++
							}
						}
					}
				}
			}
		}

		zeroes := 0
		for y := range octopi {
			for x := range octopi[y] {
				if octopi[y][x] == -1 {
					zeroes++
					octopi[y][x] = 0
				}
			}
		}

		if zeroes == 100 {
			answer = i + 1
			return
		}
	}

	return
}
