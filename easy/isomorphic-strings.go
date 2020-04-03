func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		set1 := make([]int, 256)
		set2 := make([]int, 256)
		for i := 0; i < len(s); i++ {
			val1 := int(s[i])
			val2 := int(t[i])
			if set1[val1] != set2[val2] {
				return false
			} else {
				set1[val1] = i + 1
				set2[val2] = i + 1
			}
		}
	}
	return true
}
