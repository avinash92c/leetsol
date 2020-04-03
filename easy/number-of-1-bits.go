func hammingWeight(num uint32) int {
	result := 0
	for _, val := range fmt.Sprintf("%b", num) {
		if string(val) == "1" {
			result++
		}
	}
	return result
}
