func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		lfs := addnums(nums[:i])
		rhs := addnums(nums[i+1:])
		// println(i, lfs, rhs)
		if lfs == rhs {
			return i
		}
	}
	return -1
}

func addnums(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
