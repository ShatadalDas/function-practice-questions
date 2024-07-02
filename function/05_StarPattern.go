package function


// Prints the star pattern for the given number of rows
func StarPattern(rows uint) {
	for i := uint(1); i <= rows; i++ {
		for j := uint(1); j <= i; j++ {
			print("* ")
		}
		println()
	}
}