func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 0 {
		mid := len(nums1) / 2
		return float64(nums1[mid]+nums1[mid-1]) / 2
	} else {
		mid := len(nums1) / 2
		return float64(nums1[mid])
	}
}
