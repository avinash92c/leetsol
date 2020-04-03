func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			result[i] = []int{1}
		} else if i == 1 {
			result[i] = []int{1, 1}
		} else {
			subarr := []int{1, 1}
			prevarr := result[i-1]
			for j := 1; j < len(prevarr); j++ {
				val := prevarr[j] + prevarr[j-1]
				subarr = append(subarr, 0)     // Step 1
				copy(subarr[j+1:], subarr[j:]) // Step 2
				subarr[j] = val
			}
			result[i] = subarr
		}
	}
	return result
}
