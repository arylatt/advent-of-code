package aoc202103

import (
	"strconv"
	"strings"
)

func Part1(input string) (output string) {
	lines := strings.Split(input, "\r\n")
	lines = lines[:len(lines)-1]

	gammaRate := ""

	length := len(lines[0]) - 1

	for i := 0; i <= length; i++ {
		highBits, lowBits := 0, 0

		for _, line := range lines {

			bytes := strings.Split(line, "")
			if bytes[i] == "0" {
				lowBits++
			} else {
				highBits++
			}
		}

		if highBits > lowBits {
			gammaRate += "1"
		} else {
			gammaRate += "0"
		}
	}

	epsilonRate := ""

	for _, bit := range strings.Split(gammaRate, "") {
		if bit == "0" {
			epsilonRate += "1"
		} else {
			epsilonRate += "0"
		}
	}

	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 0)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 0)

	return strconv.Itoa(int(gammaRateInt) * int(epsilonRateInt))
}

func Part2(input string) (output string) {
	lines := strings.Split(input, "\r\n")
	lines = lines[:len(lines)-1]

	oxyGenRating, co2ScrubRating := "", ""
	length := len(lines[0]) - 1

	for i := 0; i <= length; i++ {
		highBits, lowBits := 0, 0

		if len(lines) == 1 {
			oxyGenRating = lines[0]
			break
		}

		for _, line := range lines {
			bytes := strings.Split(line, "")
			if bytes[i] == "0" {
				lowBits++
			} else {
				highBits++
			}
		}

		if highBits >= lowBits {
			oxyGenRating += "1"
		} else {
			oxyGenRating += "0"
		}

		ret := []string{}
		for _, line := range lines {
			if strings.HasPrefix(line, oxyGenRating) {
				ret = append(ret, line)
			}
		}
		lines = ret
	}

	lines = strings.Split(input, "\r\n")
	lines = lines[:len(lines)-1]
	for i := 0; i <= length; i++ {
		highBits, lowBits := 0, 0

		if len(lines) == 1 {
			co2ScrubRating = lines[0]
			break
		}

		for _, line := range lines {
			bytes := strings.Split(line, "")
			if bytes[i] == "0" {
				lowBits++
			} else {
				highBits++
			}
		}

		if highBits >= lowBits {
			co2ScrubRating += "0"
		} else {
			co2ScrubRating += "1"
		}

		ret := []string{}
		for _, line := range lines {
			if strings.HasPrefix(line, co2ScrubRating) {
				ret = append(ret, line)
			}
		}
		lines = ret
	}

	oxyGenRatingInt, _ := strconv.ParseInt(oxyGenRating, 2, 0)
	co2ScrubRatingInt, _ := strconv.ParseInt(co2ScrubRating, 2, 0)

	return strconv.Itoa(int(oxyGenRatingInt) * int(co2ScrubRatingInt))
}
