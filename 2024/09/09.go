package aoc202409

import (
	"slices"
	"strconv"

	"github.com/arylatt/advent-of-code/elves"
)

type Block struct {
	ID     int
	Length int
}

func Part1(input string) (output string) {
	diskMap := inputToDiskMap(input)

	diskMap = defrag(diskMap)

	return generateChecksum(diskMap)
}

func inputToDiskMap(input string) []int {
	diskMap := []int{}

	id := 0
	freeSpace := false

	for _, char := range input {
		for range elves.Atoi(string(char)) {
			if freeSpace {
				diskMap = append(diskMap, -1)
			} else {
				diskMap = append(diskMap, id)
			}
		}

		if !freeSpace {
			id++
		}

		freeSpace = !freeSpace
	}

	return diskMap
}

func defrag(diskMap []int) []int {
	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] == -1 {
			continue
		}

		firstFree := slices.Index(diskMap, -1)
		if firstFree == -1 {
			break
		}

		if firstFree > i {
			break
		}

		diskMap[firstFree] = diskMap[i]
		diskMap[i] = -1
	}

	return diskMap
}

func generateChecksum(diskMap []int) string {
	answer := 0

	for i, id := range diskMap {
		if id == -1 {
			continue
		}

		answer += i * id
	}

	return strconv.Itoa(answer)
}

func Part2(input string) (output string) {
	diskMapBlocks := inputToBlockDiskMap(input)

	diskMapBlocks = defragBlocks(diskMapBlocks)

	diskMap := blockDiskMapToInts(diskMapBlocks)

	return generateChecksum(diskMap)
}

func inputToBlockDiskMap(input string) []Block {
	diskMap := []Block{}

	id := 0
	freeSpace := false

	for _, char := range input {
		block := Block{Length: elves.Atoi(string(char))}
		if freeSpace {
			block.ID = -1
		} else {
			block.ID = id
		}

		diskMap = append(diskMap, block)

		if !freeSpace {
			id++
		}

		freeSpace = !freeSpace
	}

	return diskMap
}

func defragBlocks(diskMap []Block) []Block {
	i := len(diskMap) - 1

	for i > 0 {
		if diskMap[i].ID == -1 {
			i--
			continue
		}

		firstFree := slices.IndexFunc(diskMap, func(b Block) bool {
			return b.ID == -1 && b.Length >= diskMap[i].Length
		})

		if firstFree == -1 || firstFree > i {
			i--
			continue
		}

		freeBlock := diskMap[firstFree]

		diskMap[firstFree] = diskMap[i]
		diskMap[i] = Block{ID: -1, Length: diskMap[i].Length}

		if freeBlock.Length > diskMap[firstFree].Length {
			diskMap = slices.Insert(diskMap, firstFree+1, Block{ID: -1, Length: freeBlock.Length - diskMap[firstFree].Length})
			i++
		}

		i--
	}

	return diskMap
}

func blockDiskMapToInts(diskMap []Block) []int {
	ints := []int{}

	for _, block := range diskMap {
		for range block.Length {
			ints = append(ints, block.ID)
		}
	}

	return ints
}
