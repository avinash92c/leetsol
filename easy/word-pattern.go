func wordPattern(pattern string, str string) bool {
	if len(pattern) == 0 || len(str) == 0 {
		return false
	}
	m := make(map[string]string)
	v := strings.Split(str, " ")
	if len(pattern) != len(v) {
		return false
	}
	for i, p := range pattern {
		fmt.Println(m[string(p)], v[i])
		if len(m[string(p)]) == 0 {
			m[string(p)] = v[i]
		} else if m[string(p)] != v[i] {
			return false
		}
	}
	prev := ""
	for _, v := range m {
		if prev == v {
			return false
		}
		prev = v
	}
	return true
}
