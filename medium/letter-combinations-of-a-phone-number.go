func letterCombinations(digits string) []string {
	mappings := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result := make([]string, 0)

	for t := 0; t < len(digits); t++ {
		pos, _ := strconv.Atoi(string(digits[t]))
		if len(result) == 0 {
			result = append(result, strings.Split(mappings[pos-2], "")...)
		} else {
			result = mergeSlices(result, strings.Split(mappings[pos-2], ""))
		}
	}
	return result
}

func mergeSlices(a1, a2 []string) []string {
	var result []string
	for _, v1 := range a1 {
		for _, v2 := range a2 {
			result = append(result, v1+v2)
		}
	}
	return result
}
