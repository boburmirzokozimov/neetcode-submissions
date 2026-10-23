func maxProfit(prices []int) int {
	profit := 0
	l, r := 0, 1

	for r < len(prices) {
		if prices[r] < prices[l] {
			l = r
		} else {
			profit = max(profit, prices[r] - prices[l])
		}
		r++
	}

//[10,1,5,6,7,1] => -9
	return profit
}
