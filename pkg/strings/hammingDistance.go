package strings

func HammingDistance(a, b string) int {
	if len(a) != len(b) {
		panic("Hamming distance only defined for strings of equal length")
	}

	distance := 0

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance
}
