package easy

/* 1678. Goal Parser Interpretation */
func interpret(command string) string {
	var result string

	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			result += `G`
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				result += `o`
				i++
			} else {
				if i+2 < len(command) && i+3 < len(command) {
					if command[i+2] == 'l' && command[i+3] == ')' {
						result += `al`
						i += 3
					}
				}
			}
		}
	}

	return result
}
