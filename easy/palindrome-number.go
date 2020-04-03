func isPalindrome(x int) bool {
	input := strconv.Itoa(x)
	if input == reverseString(input) {
		return true
	}
	return false
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
