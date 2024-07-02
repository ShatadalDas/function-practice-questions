package function


// Calculates the factorial of a number.
// If the number is negative, returns -1.
// If the number is greater than 20, returns 0, means the function is too complex to calculate.
func Factorial(n int) int {
	switch {
		case n > 20:
			return 0
		case n < 0:
			return -1
		case n == 0:
			return 1
		case n <= 2:
			return n
	}

	fact := 1
	for i := range n {
		i++
		fact *= i
	}

	return fact
}