func maxProfit(prices []int) int {
	profit := 0
	if len(prices) == 0 {
return profit
	}
	min := prices[0]

	for _, val := range prices {
		fmt.Println(min, val)
		if val > min {
			profit += val - min
		}
		min = val
	}

	return profit
}
