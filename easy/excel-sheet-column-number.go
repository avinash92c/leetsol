func titleToNumber(s string) int {
	initial := 65
	sol := 0
	pos := 0
	for i := len(s) - 1; i >= 0; i-- {
		sol += (((int(s[i]) - initial) + 1) * int(math.Pow(float64(26), float64(pos))))
		pos++
	}
	return sol
}
