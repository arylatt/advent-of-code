package elves

import "strings"

func SplitIntoLines(input string) (ret []string) {
	inputs := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")

	for _, input := range inputs {
		if strings.TrimSpace(input) != "" {
			ret = append(ret, input)
		}
	}

	return
}
