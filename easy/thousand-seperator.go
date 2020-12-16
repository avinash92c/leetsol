package easy

import (
	"fmt"
	"strconv"
)

/* 1556. Thousand Separator */
func thousandSeparator(n int) string {
	result := ``
	nstr := strconv.Itoa(n)
	nlen := len(nstr)

	if n < 1000 {
		return strconv.Itoa(n)
	}

	start, end := nlen-3, nlen

	//strip every three from end and append
	for {
		fmt.Println(result, start, end)
		result = nstr[start:end] + result
		end = start
		start = start - 3
		if start <= 0 {
			start = 0
			result = `.` + result
			result = nstr[start:end] + result
			break
		} else {
			result = `.` + result
		}
	}

	return result
}
