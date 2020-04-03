func convert(s string, numRows int) string {
	if len(s) == numRows || numRows == 1 {
		// fmt.Println(s)
		return s
	}
	resarr := make([]string, numRows)
	flip := false
	j := 0
	for i := 0; i < len(s); i++ {
		// fmt.Println(j, flip)
		resarr[j] += string(s[i])
		if flip {
			j--
		} else {
			j++
		}

		if j < 0 {
			flip = !flip
			j += 2
		}
		if j == numRows-1 {
			flip = !flip
		}
	}
	var result string
	for _, v := range resarr {
		result += v
	}
	// fmt.Println(result)
	return result
}
