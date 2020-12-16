package easy

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1, s2 strings.Builder

	for _, val := range word1 {
		s1.WriteString(val)
	}

	for _, val := range word2 {
		s2.WriteString(val)
	}

	if s1.String() == s2.String() {
		return true
	}

	return false
}
