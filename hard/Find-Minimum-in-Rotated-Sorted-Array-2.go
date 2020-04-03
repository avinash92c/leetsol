func findMin(nums []int) int {
	min := nums[0]
	if len(nums) == 1{
		return min
	}
	for i:=1; i<len(nums);i++{
		if min> nums[i]{
			return nums[i]
		}
	}
	return min
}
