func convertToTitle(n int) string {
	initial := 65 //FOR CHARACTER CONVERSION
	result := ""
	for ; n > 0; n = (n - 1) / 26 {
		chtag := (n - 1) % 26
		result = string(initial + chtag)+result
	}
	return result
}
