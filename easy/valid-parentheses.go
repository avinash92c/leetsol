func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	controlstruct := make([]string, 0)
	for i := range s {
		if len(controlstruct) > 0 && mapping[string(s[i])] == controlstruct[len(controlstruct)-1] {
			controlstruct = controlstruct[:len(controlstruct)-1]
		} else {
			controlstruct = append(controlstruct, string(s[i]))
		}
	}

	if len(controlstruct) > 0 {
		return false
	}
	return true
}

var mapping map[string]string = map[string]string{"}": "{", "]": "[", ")": "("}
