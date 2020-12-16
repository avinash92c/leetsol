package easy

func decrypt(code []int, k int) []int {
	codelen := len(code)
	result := make([]int, codelen)

	if k == 0 {
		return result
	} else {
		for i := 0; i < codelen; i++ {
			if k > 0 {
				result[i] = getksum(code, i+1, k)
			} else {
				result[i] = getksum(code, i-1, k)
			}
		}
	}

	return result
}

func getksum(code []int, i, k int) int {
	var result int
	l := k
	if k < 0 {
		l = k * -1
	}

	for j := 0; j < l; j++ {
		if len(code) == i {
			i = 0
		} else if i < 0 {
			i = len(code) - 1
		}

		result += code[i]

		if k < 0 {
			i--
		} else if k > 0 {
			i++
		}
	}

	return result
}
