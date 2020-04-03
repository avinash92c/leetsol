
func findMin(nums []int) int {
	max := len(nums)
    min := nums[0] 
	for i := range nums {
		if i+1 < max  && nums[i] > nums[i+1] {
			min = nums[i+1]
            break
		}
	}
    return min
}
