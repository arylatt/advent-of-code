package aoc201504

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func FindMD5IterationWithPrefix(start int, hashInput, targetPrefix string) (result int) {
	result = start
	for {
		b := (md5.Sum([]byte(fmt.Sprintf("%s%d", hashInput, result))))
		if hex.EncodeToString(b[:])[0:5] == targetPrefix {
			return result
		}
		result++
	}
}

func Part1(input string) (output string) {
	return strconv.Itoa(FindMD5IterationWithPrefix(1, input, "00000"))
}

func Part2(input string) (output string) {
	return strconv.Itoa(FindMD5IterationWithPrefix(254575, input, "000000"))
}
