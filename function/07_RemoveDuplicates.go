package function


// Takes an array of integers and returns a new array with no duplicates
// and the order of the elements in the new array is the same as the original one
func RemoveDuplicates(arr []int) []int {
	seen := map[int]bool{}
	res := []int{}
	
	for _, el := range arr {
		if !seen[el] {
			seen[el] = true
			res = append(res, el)
		}
	}

	return res
}
