package easy

/* 1672. Richest Customer Wealth */
func maximumWealth(accounts [][]int) int {
	result := 0

	for _, caccounts := range accounts {
		wealth := 0
		for _, cwealth := range caccounts {
			wealth += cwealth
		}
		if result < wealth {
			result = wealth
		}
	}

	return result
}
