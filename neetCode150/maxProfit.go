package neetcode150

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func MaxProfit(prices []int) int {
	length := len(prices)
	if length == 0 || length == 1 {
		return 0
	}
	buyingPrice := prices[0]
	sellingPrice := 0
	profit := 0
	maxProfit := 0
	for i := 1; i < length; i++ {
		if prices[i] < buyingPrice {
			buyingPrice = prices[i]
			sellingPrice = 0
		} else {
			if prices[i] > sellingPrice {
				sellingPrice = prices[i]
			}
		}
		profit = sellingPrice - buyingPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
