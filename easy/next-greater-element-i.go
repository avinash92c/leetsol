
func prepScanMap(nums []int) map[int]int {
	smap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		flip := true
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				smap[nums[i]] = nums[j]
				flip = false
				break
			}
		}

		if flip {
			smap[nums[i]] = -1
		}
	}
	return smap
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	smap := prepScanMap(nums2)
	for _, val := range nums1 {
		result = append(result, smap[val])
	}
	return result
}
