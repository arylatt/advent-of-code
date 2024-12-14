package aoc202412

import (
	"slices"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

type Region struct {
	ID    int
	Plant string
	Plots []elves.Point
}

func (r Region) Adjacent(p elves.Point) bool {
	for _, plot := range r.Plots {
		if plot.X == p.X && plot.Y == p.Y-1 {
			return true
		}
		if plot.X == p.X && plot.Y == p.Y+1 {
			return true
		}
		if plot.X == p.X-1 && plot.Y == p.Y {
			return true
		}
		if plot.X == p.X+1 && plot.Y == p.Y {
			return true
		}
	}

	return false
}

func (r Region) Contains(p elves.Point) bool {
	for _, plot := range r.Plots {
		if plot.X == p.X && plot.Y == p.Y {
			return true
		}
	}

	return false
}

func Part1(input string) (output string) {
	grid := elves.SplitIntoLines(input)
	regions := findAllRegions(grid)

	// calculate perimeters and total the costs
	answer := 0

	for _, region := range regions {
		perimeter := 0

		for _, plot := range region.Plots {
			if !region.Contains(plot.Shift(0, -1)) {
				perimeter++
			}
			if !region.Contains(plot.Shift(0, 1)) {
				perimeter++
			}
			if !region.Contains(plot.Shift(-1, 0)) {
				perimeter++
			}
			if !region.Contains(plot.Shift(1, 0)) {
				perimeter++
			}
		}

		answer += perimeter * len(region.Plots)
	}

	return strconv.Itoa(answer)
}

func findRegions(regions []*Region, plant string) (matching []*Region) {
	for _, region := range regions {
		if region.Plant == plant {
			matching = append(matching, region)
		}
	}
	return
}

func findAllRegions(grid []string) (regions []*Region) {

	regionID := 0

	// find all the plots
	for y, row := range grid {
		for x, cell := range row {
			plant := string(cell)
			plot := elves.Point{X: x, Y: y}

			matchingRegions := findRegions(regions, plant)

			if len(matchingRegions) == 0 {
				regions = append(regions, &Region{
					ID:    regionID,
					Plant: plant,
					Plots: []elves.Point{plot},
				})

				regionID++
			} else {
				appended := false

				for _, region := range matchingRegions {
					if region.Adjacent(plot) {
						region.Plots = append(region.Plots, plot)
						appended = true
						break
					}
				}

				if !appended {
					regions = append(regions, &Region{
						ID:    regionID,
						Plant: plant,
						Plots: []elves.Point{plot},
					})

					regionID++
				}
			}
		}
	}

	// deduplicate plots and merge them together
	for i := 0; i < len(regions); i++ {
		for j := 0; j < len(regions); j++ {
			if i == j {
				continue
			}

			if regions[i].Plant != regions[j].Plant {
				continue
			}

			if regions[i].ID == regions[j].ID {
				continue
			}

			for _, plot := range regions[j].Plots {
				if regions[i].Adjacent(plot) {
					regions[i].Plots = append(regions[i].Plots, regions[j].Plots...)
					regions = slices.Delete(regions, j, j+1)
					j--
					break
				}
			}
		}
	}

	return
}

func Part2(input string) (output string) {
	grid := elves.SplitIntoLines(input)
	regions := findAllRegions(grid)

	// calculate perimeters and total the costs
	answer := 0

	for _, region := range regions {
		corners := 0

		for _, plot := range region.Plots {
			if !region.Contains(plot.Shift(0, -1)) && !region.Contains(plot.Shift(-1, 0)) {
				corners++
			}
			if !region.Contains(plot.Shift(0, -1)) && !region.Contains(plot.Shift(1, 0)) {
				corners++
			}
			if !region.Contains(plot.Shift(1, 0)) && !region.Contains(plot.Shift(0, 1)) {
				corners++
			}
			if !region.Contains(plot.Shift(0, 1)) && !region.Contains(plot.Shift(-1, 0)) {
				corners++
			}

			if region.Contains(plot.Shift(0, -1)) && region.Contains(plot.Shift(-1, 0)) && !region.Contains(plot.Shift(-1, -1)) {
				corners++
			}
			if region.Contains(plot.Shift(0, -1)) && region.Contains(plot.Shift(1, 0)) && !region.Contains(plot.Shift(1, -1)) {
				corners++
			}
			if region.Contains(plot.Shift(1, 0)) && region.Contains(plot.Shift(0, 1)) && !region.Contains(plot.Shift(1, 1)) {
				corners++
			}
			if region.Contains(plot.Shift(0, 1)) && region.Contains(plot.Shift(-1, 0)) && !region.Contains(plot.Shift(-1, 1)) {
				corners++
			}
		}

		answer += corners * len(region.Plots)
	}

	return strconv.Itoa(answer)
}
