func reverse(n int) int {
	var negative bool
	if n < 0 {
		negative = true
		n = n * -1
	}

	new_int := 0
	for n > 0 {
		new_int = new_int*10 + n%10
		n /= 10
	}

	if new_int > math.MaxInt32 || new_int < math.MinInt32 {
		return 0
	}

	if negative {
		return new_int * -1
	} else {
		return new_int
	}
}
