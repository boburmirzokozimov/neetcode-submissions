func maxProfit(prices []int) int {
	profit := 0
	l, r := 0, 1

	for r < len(prices) {
		if prices[r] > prices[l] {
			profit = max(profit, prices[r] - prices[l])
		} else {
			l = r
		}
		r++
	}

//[10,1,5,6,7,1] => -9
	return profit
}
