func isMatch(s string, p string) bool {
    str := strings.Trim(p, " ")
    result, _ := regexp.MatchString(fmt.Sprintf("^%s$",str), s)
	return result
}
