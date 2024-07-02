package function

// Calculates the number of permutations of n objects taken r at a time.
// If n < 0 or r < 0 or r > n, returns -1, which can be interpreted as "invalid input".
func Permutation(n int, r int) int {
	if n < 0 || r < 0 || n < r {
		return -1
	}

	p := 1

	for i := n; i > n - r; i-- {
		p *= i
	}

	return p
}
