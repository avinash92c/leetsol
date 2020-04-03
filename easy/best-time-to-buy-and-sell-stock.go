func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				profitl := prices[j] - prices[i]
				if profit < profitl {
					profit = profitl
				}
			}
		}
	}
	return profit
}
