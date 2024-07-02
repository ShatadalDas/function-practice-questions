package function

func Combination(n int, r int) int {
	if n < 0 || r < 0 || n < r {
		return -1
	}

	c := 1

	maxi := max(n - r, r)
	mini := min(n - r, r)

	for i := n; i > maxi; i-- {
		c *= i
	}

	c /= Factorial(mini)

	return c
}