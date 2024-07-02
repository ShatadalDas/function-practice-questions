package function


// Calculates the sum of the digits of a number.
// If the number is negative, the digits are calculated from the absolute value of the number.
func SomeOfDigits(n int) int {
    if n < 0 {
			n *= -1
		}

		sum := 0

		for n > 0 {
			lastDigit := n % 10
			sum += lastDigit
			n /= 10
		}

		return sum
}