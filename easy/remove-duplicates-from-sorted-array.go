
func removeDuplicates(nums []int) int {
    if nums == nil|| len(nums)==0 {
		return 0
	}
	prevval := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prevval {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			prevval = nums[i]
		}
	}
	return len(nums)
}
