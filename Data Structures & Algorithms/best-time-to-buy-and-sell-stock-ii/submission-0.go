
func maxProfit(prices []int) int {
    profit := 0

    for i := 1; i < len(prices); i++ {
		prev := prices[i-1]
		if prev < prices[i] {
			profit += prices[i] - prev
		}
    }

    return profit
}