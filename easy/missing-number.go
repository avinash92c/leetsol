func missingNumber(nums []int) int {
	n := len(nums)
	expectedsum, actualsum := (n*(n+1))/2, 0
	for _, val := range nums {
		actualsum += val
	}
	return expectedsum - actualsum
}
