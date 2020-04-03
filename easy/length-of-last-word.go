func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if strings.Index(s, " ") < 0 {
		return len(s)
	} else {
		r := strings.Split(s, " ")
		return len(r[len(r)-1])
	}
}
