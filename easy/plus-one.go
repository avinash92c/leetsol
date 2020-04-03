func plusOne(digits []int) []int {

	if digits[len(digits)-1] == 9 {
		for i := len(digits) - 1; i >= 0; i-- {
			if i == 0 && digits[i] == 9 {
				digits[i] = 0
				return append([]int{1}, digits...)
			} else if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i] += 1
				break
			}
		}
	} else {
		digits[len(digits)-1]++
	}
	return digits
}
