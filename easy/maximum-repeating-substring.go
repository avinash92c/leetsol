package easy

/* 1668. Maximum Repeating Substring */
func maxRepeating(sequence string, word string) int {
	var result int

	//SLIDING SCAN
	wordlen := len(word)
	sequencelen := len(sequence)

	if sequencelen == wordlen && sequence == word {
		return 1
	}

	start := 0
	end := start + wordlen
	count := 0

	for end <= sequencelen {
		if sequence[start:end] == word {
			count++
			start = end
		} else {
			count = 0
			start++
		}
		end = start + wordlen

		if count > result {
			result = count
		}
	}

	return result
}

/*
func comparestrings(s1, s2 string) (int, bool) {
	if s1 != s2 {
		for i := 0; i < len(s2); i++ {
			if s1[i] != s2[i] {
				return i, false
			}
		}
	}
	return -1, true
}
*/
/*
func maxRepeating(sequence string, word string) int {
	var result int
	carr := make([]int, 26)

	for _, ch := range word {
		carr[ch-'a'] = 1
	}

	count := 0
	seqlen := len(sequence)
	//scan sequence
	for i := 0; i < seqlen; {

		match := true
		for _, val := range word {
			if i < seqlen && byte(val) != sequence[i] {
				match = false
				i++
				break
			}
			i++
		}

		if match {
			count++
		} else {
			count = 0
		}

		if result < count {
			result = count
		}
	}

	return result
}
*/
