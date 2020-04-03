// THIS IS SLOW FUNCTION
// I'LL WRITE A RABIN-KARP SOLUTION AND POST LATER
func isSubsequence(s string, t string) bool {
   if len(strings.Trim(s, " ")) == 0 {
		return true
	}
	if len(strings.Trim(t, " ")) == 0 {
		return false
	}
    // fmt.Println(len(t))
	sidx := 0
	tidx := 0
	for tidx < len(t) {
        // fmt.Println(tidx)
		if string(t[tidx]) == string(s[sidx]) {
			sidx++
		}
        if sidx == len(s) {
			break
		}
		tidx++
	}
	return sidx == len(s)
}
