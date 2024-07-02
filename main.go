package main

import (
	"fmt"
	"function-practice-questions/function"
)

func main() {
	f1 := function.Factorial(5)
	f2 := function.SomeOfDigits(555)
	f3 := function.Permutation(5, 2)
	f4 := function.Combination(5, 2)

	println(f1)
	println(f2)
	println(f3)
	println(f4)

	function.StarPattern(5)
	function.PascalTriangle(5)
	res := function.RemoveDuplicates([]int{1, 1, 2, 2, 3, 1, 7, 1, 3, 2, 3})
	fmt.Println(res)
}
