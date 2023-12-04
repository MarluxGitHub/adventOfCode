package math

func Union(a, b []int) []int {
	m := make(map[int]bool)
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		m[item] = true
	}

	var result []int
	for item := range m {
		result = append(result, item)
	}

	return result
}

func Intersection(a, b []int) []int {
	m := make(map[int]bool)
	for _, item := range a {
		m[item] = true
	}

	var result []int
	for _, item := range b {
		if m[item] {
			result = append(result, item)
		}
	}

	return result
}

func Difference(a, b []int) []int {
	m := make(map[int]bool)
	for _, item := range a {
		m[item] = true
	}

	var result []int
	for _, item := range b {
		if !m[item] {
			result = append(result, item)
		}
	}

	return result
}