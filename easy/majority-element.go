func majorityElement(nums []int) int {
	sort.Ints(nums)
	currnum := 0
	// prevnum := 0
	count := 0
	// found := false
	n := len(nums)

	for _, num := range nums {
		if currnum == num {
			count++
		} else {
			// prevnum = currnum
			// fmt.Println(prevnum)
			if count > n/2 {
				break
			} else {
				currnum = num
				count = 1
			}
		}
	}
	return currnum
}
