package interview

import (
	"strconv"
	"strings"
)

func HighestEarning(s string) int {
	inputs := strings.Split(
		strings.Replace(s, " ", "", -1),
		",",
	)

	earning := 0
	highest := 0

	for _, input := range inputs {
		if i, err := strconv.Atoi(input); err == nil {
			earning = earning + i

			if earning < 0 { // earning can't be less than 0
				earning = 0
			}
		} else {
			return 0
		}

		if earning > highest {
			highest = earning
		}
	}

	return highest
}
