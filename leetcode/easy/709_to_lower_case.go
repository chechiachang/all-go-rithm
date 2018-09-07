package easy

func toLowerCase(str string) string {
	lower := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
		"D": "d",
		"E": "e",
		"F": "f",
		"G": "g",
		"H": "h",
		"I": "i",
		"J": "j",
		"K": "k",
		"L": "l",
		"M": "m",
		"N": "n",
		"O": "o",
		"P": "p",
		"Q": "q",
		"R": "r",
		"S": "s",
		"T": "t",
		"U": "u",
		"V": "v",
		"W": "w",
		"X": "x",
		"Y": "y",
		"Z": "z",
	}

	result := ""
	for i := 0; i < len(str); i++ {

		// str[i] is capital (has key in lower), map to lower case
		if val, ok := lower[string(str[i])]; ok {
			result += val

			// str[i] is lower case, append result
		} else {
			result += string(str[i])
		}

	}
	return result
}
