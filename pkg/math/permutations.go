package math

func Permutations(n int) [][]int {
	perms := make([][]int, 0)
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}

	// generate all permutations
	permute(&perms, perm, 0, n)

	return perms
}

func permute(perms *[][]int, perm []int, i int, n int) {
	if i == n {
		permCopy := make([]int, len(perm))
		copy(permCopy, perm)
		*perms = append(*perms, permCopy)
		return
	}

	for j := i; j < n; j++ {
		perm[i], perm[j] = perm[j], perm[i]
		permute(perms, perm, i+1, n)
		perm[i], perm[j] = perm[j], perm[i]
	}
}
