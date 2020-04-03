func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
		i = i + 1
	}
	return nums[len(nums)-1]
}
