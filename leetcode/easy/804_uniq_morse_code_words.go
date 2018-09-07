package easy

func uniqueMorseRepresentations(words []string) int {
	morse := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}
	set := map[string]bool{}

	for _, w := range words {
		morsedW := ""

		for i := 0; i < len(w); i++ {
			morsedW = morsedW + morse[string(w[i])]
		}

		set[morsedW] = true
	}

	return len(set)
}
