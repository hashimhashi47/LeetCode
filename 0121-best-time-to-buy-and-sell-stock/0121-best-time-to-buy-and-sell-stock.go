func maxProfit(prices []int) int {
	buy := prices[0]
	sell := 0
	for i := 0; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		}
		if sell < prices[i]-buy {
			sell = prices[i] - buy
		}
	}
	return sell
}