package easy

/*1684. Count the Number of Consistent Strings */
func countConsistentStrings(allowed string, words []string) int {
	count := 0
	allowedchars := make([]int, 26) //only for small chars
	for _, alchar := range allowed {
		allowedchars[alchar-'a'] = 1
	}

	//SCAN WORDS
	for _, word := range words {
		match := true
		for _, cchar := range word {
			if allowedchars[cchar-'a'] != 1 {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}

/*
func countConsistentStrings(allowed string, words []string) int {
	count := 0
	regexpat := fmt.Sprintf(`^[%s]+$`, allowed)
	for _, word := range words {
		matched, _ := regexp.MatchString(regexpat, word)
		fmt.Println(matched) // false
		if matched {
			count++
		}
	}

	return count
}
*/
/*
func countConsistentStrings(allowed string, words []string) int {
	count := 0
	allowarr := strings.Split(allowed, ``)
	allowmap := make(map[string]bool)
	for _, allowchar := range allowarr {
		allowmap[allowchar] = true
	}

	//EVALUATE
	for _, word := range words {
		wordarr := strings.Split(word, ``)
		clean := true
		for _, wordchar := range wordarr {
			if _, ok := allowmap[wordchar]; !ok {
				clean = false
			}
		}
		if clean {
			count++
		}
	}
	return count
}
*/
