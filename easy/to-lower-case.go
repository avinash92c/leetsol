func toLowerCase(str string) string {
    var res string
	for _, val := range str {
		if val > 64 && val < 91 {
			val += 32
		}
		res += string(val)
	}
	return res
}
