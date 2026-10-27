func maxProfit(prices []int) int {
	maxProfit := 0
	buy := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[buy] {
			buy = i
		} else {
			maxProfit = max(maxProfit, prices[i] - prices[buy])
		}
	}

	return maxProfit
}
