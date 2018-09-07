package easy

func judgeCircle(moves string) bool {
	up := 0
	right := 0

	for i := 0; i < len(moves); i++ {
		switch string(moves[i]) {
		case "U":
			up++
		case "D":
			up--
		case "R":
			right++
		case "L":
			right--
		}
	}

	return up == 0 && right == 0
}
