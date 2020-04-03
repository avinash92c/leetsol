func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sort.Ints(nums)
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		if num == nums[i] {
			return true
		}
		num = nums[i]
	}
	return false
}
