package function


// Prints the pascal triangle for the given number of rows
func PascalTriangle(rows uint) {
	prevRow := make([]uint, rows+1)

	// the first row
	prevRow[1] = 1
	println(prevRow[1])

	// 2 to rows
	for i := uint(2); i <= rows; i++ {
		prev := uint(1)
		
		// 1 to i
		for j := uint(1); j <= i; j++ {

			curr := prevRow[j] + prevRow[j-1]
			print(curr, " ")

			if j > 1 {
				prevRow[j-1] = prev
				prev = curr
			}
		}
		prevRow[i] = prev

		println()
	}

}
