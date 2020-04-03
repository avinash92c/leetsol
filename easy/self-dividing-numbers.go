func selfDividingNumbers(left int, right int) []int {
	var nums []int
	for i := left; i <= right; i++ {
		// hz, digits := haszero(i)
		// if !hz {
		skip := false
		// for _, val := range digits {
		num := i
		for num != 0 {
			// println(val, i, i%val)
			digit := num % 10
			if digit == 0 || i%digit != 0 {
				// if i%val != 0 {
				skip = true
				break
			}
			num /= 10
			// println(num)
		}
		if !skip {
			nums = append(nums, i)
		}
		// }
	}
	return nums
}
