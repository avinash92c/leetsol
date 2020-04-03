func myAtoi(str string) int {
	result := 0
	number := strings.Trim(str, " ") //REMOVED START AND END WHITESPACES
	if len(number) == 0 {
		return 0
	}
	sign := false
	i := 0
	//CHECK OF SIGN
	if int(number[0]) == 45 {
		sign = true
		i = 1
	} else if int(number[0]) == 43 {
		i = 1
	}
	for ; i < len(number); i++ {
		val := number[i]
		if int(val) > 47 && int(val) < 58 {
			result = result*10 + (int(val) - int('0'))
		} else {
			break
		}
		resval := result
		if sign {
			resval *= -1
		}
		if resval > math.MaxInt32 {
			return math.MaxInt32
		}
		if resval < math.MinInt32 {
			return math.MinInt32
		}
	}

	if sign {
		result *= -1
	}

	return result
}
