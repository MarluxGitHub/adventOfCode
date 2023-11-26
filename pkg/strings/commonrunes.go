package strings

func GetMapOfRuneCounts(s string) map[rune]int {
	m := make(map[rune]int)

	for _, r := range s {
		m[r]++
	}

	return m
}

func GetMostCommonRune(s string) rune {
	m := GetMapOfRuneCounts(s)

	max := 0
	var maxRune rune

	for r, count := range m {
		if count > max {
			max = count
			maxRune = r
		}
	}

	return maxRune
}

func GetLeastCommonRune(s string) rune {
	m := GetMapOfRuneCounts(s)

	min := 9999999
	var minRune rune

	for r, count := range m {
		if count < min {
			min = count
			minRune = r
		}
	}

	return minRune
}
