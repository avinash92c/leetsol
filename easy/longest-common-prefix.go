func longestCommonPrefix(strs []string) string {
	longest := ""
	if len(strs) != 0 {
		longest = strs[0]
	}
	for i := 1; i < len(strs); i++ {
        fmt.Println("str-"+strs[i])
		for !strings.HasPrefix(strs[i], longest) {
			longest = longest[0 : len(longest)-1]
		}
	}
	return longest
}
