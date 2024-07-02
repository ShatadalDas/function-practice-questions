package function_test

import (
	"testing"
	"function-practice-questions/function"
)

func TestFactorial(t *testing.T) {
	if function.Factorial(5) != 120 {
		t.Errorf("Factorial(5) should be 120, but is %d", function.Factorial(5))
	}
	if function.Factorial(10) != 3628800 {
		t.Errorf("Factorial(10) should be 3628800, but is %d", function.Factorial(10))
	}
	if function.Factorial(0) != 1 {
		t.Errorf("Factorial(0) should be 1, but is %d", function.Factorial(0))
	}
	if function.Factorial(-834) != -1 {
		t.Errorf("Factorial(-834) should be -1, but is %d", function.Factorial(-834))
	}
	if function.Factorial(34) != 0 {
		t.Errorf("Factorial(34) should be 0, but is %d", function.Factorial(34))
	}
}

func TestSomeOfDigits(t *testing.T) {
	if function.SomeOfDigits(12345) != 15 {
		t.Errorf("SomeOfDigits(12345) should be 15, but is %d", function.SomeOfDigits(12345))
	}
	if function.SomeOfDigits(0) != 0 {
		t.Errorf("SomeOfDigits(0) should be 0, but is %d", function.SomeOfDigits(0))
	}
	if function.SomeOfDigits(-12345) != 15 {
		t.Errorf("SomeOfDigits(-12345) should be 15, but is %d", function.SomeOfDigits(-12345))
	}
}

func TestPermutation(t *testing.T) {
	if function.Permutation(5, 2) != 20 {
		t.Errorf("Permutation(5, 2) should be 20, but is %d", function.Permutation(5, 2))
	}
	if function.Permutation(5, 3) != 60 {
		t.Errorf("Permutation(5, 3) should be 60, but is %d", function.Permutation(5, 3))
	}
	if function.Permutation(5, 4) != 120 {
		t.Errorf("Permutation(5, 4) should be 120, but is %d", function.Permutation(5, 4))
	}
	if function.Permutation(5, 5) != 120 {
		t.Errorf("Permutation(5, 5) should be 120, but is %d", function.Permutation(5, 5))
	}
	if function.Permutation(5, 6) != -1 {
		t.Errorf("Permutation(5, 6) should be -1, but is %d", function.Permutation(5, 6))
	}
	if function.Permutation(0, 0) != 1 {
		t.Errorf("Permutation(0, 0) should be 1, but is %d", function.Permutation(0, 0))
	}
	if function.Permutation(1, 0) != 1 {
		t.Errorf("Permutation(1, 0) should be 1, but is %d", function.Permutation(1, 0))
	}
	if function.Permutation(1, 1) != 1 {
		t.Errorf("Permutation(1, 1) should be 1, but is %d", function.Permutation(1, 1))
	}
}


func TestCombination(t *testing.T) {
	if function.Combination(5, 2) != 10 {
		t.Errorf("Combination(5, 2) should be 10, but is %d", function.Combination(5, 2))
	}
	if function.Combination(5, 3) != 10 {
		t.Errorf("Combination(5, 3) should be 10, but is %d", function.Combination(5, 3))
	}
	if function.Combination(5, 4) != 5 {
		t.Errorf("Combination(5, 4) should be 5, but is %d", function.Combination(5, 4))
	}
	if function.Combination(5, 5) != 1 {
		t.Errorf("Combination(5, 5) should be 1, but is %d", function.Combination(5, 5))
	}
	if function.Combination(5, 6) != -1 {
		t.Errorf("Combination(5, 6) should be -1, but is %d", function.Combination(5, 6))
	}
	if function.Combination(0, 0) != 1 {
		t.Errorf("Combination(0, 0) should be 1, but is %d", function.Combination(0, 0))
	}
	if function.Combination(1, 0) != 1 {
		t.Errorf("Combination(1, 0) should be 1, but is %d", function.Combination(1, 0))
	}
	if function.Combination(1, 1) != 1 {
		t.Errorf("Combination(1, 1) should be 1, but is %d", function.Combination(1, 1))
	}
}


func TestStarPattern(t *testing.T) {
	function.StarPattern(6)
	function.StarPattern(1)
	function.StarPattern(0)
	function.StarPattern(10)
}


func TestRemoveDuplicates(t *testing.T) {
	resArray := function.RemoveDuplicates([]int{5, 1, 2, 1, 4, 3, 7, 2, 5, 1, 0, 5, 0, 1, 0, 1, 0})
	expArray := []int{5, 1, 2, 4, 3, 7, 0}

	if len(resArray) != len(expArray) {
		t.Errorf("RemoveDuplicates([5, 1, 2, 1, 4, 3, 7, 2, 5, 1, 0, 5, 0, 1, 0, 1, 0]) should be [5, 1, 2, 4, 3, 7, 0], but is %v", resArray)
	}

	for i := 0; i < len(resArray); i++ {
		if resArray[i] != expArray[i] {
			t.Errorf("RemoveDuplicates([5, 1, 2, 1, 4, 3, 7, 2, 5, 1, 0, 5, 0, 1, 0, 1, 0]) should be [5, 1, 2, 4, 3, 7, 0], but is %v", resArray)
		}
	}



	resArray = function.RemoveDuplicates([]int{5, 1, 2, 4, 3, 7, 0})
	expArray = []int{5, 1, 2, 4, 3, 7, 0}

	if len(resArray) != len(expArray) {
		t.Errorf("RemoveDuplicates([5, 1, 2, 4, 3, 7, 0]) should be [5, 1, 2, 4, 3, 7, 0], but is %v", resArray)
	}

	for i := 0; i < len(resArray); i++ {
		if resArray[i] != expArray[i] {
			t.Errorf("RemoveDuplicates([5, 1, 2, 4, 3, 7, 0]) should be [5, 1, 2, 4, 3, 7, 0], but is %v", resArray)
		}
	}



	resArray = function.RemoveDuplicates([]int{})
	expArray = []int{}

	if len(resArray) != len(expArray) {
		t.Errorf("RemoveDuplicates([]) should be [], but is %v", resArray)
	}
}