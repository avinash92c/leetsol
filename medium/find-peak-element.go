func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] > nums[i+1] {
			return i
        }else if i+1 == len(nums){
            return i
        }
	}
	return 0
}
