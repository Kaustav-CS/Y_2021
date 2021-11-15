/******************************************************************************
Link:   https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
123. Best Time to Buy and Sell Stock III

Your input
[3,3,5,0,0,3,1,4]
Output
6
Expected
6

DATE:    2021, November 15
২৯ কার্তিক,১৪২৮

*******************************************************************************/
func maxProfit(prices []int) int {
    firstSell, secondSell := 0, 0
	firstBuy, secondBuy := math.MinInt32, math.MinInt32

	for _, price := range prices {
		secondSell = max(secondSell, secondBuy+price)
		secondBuy = max(secondBuy, firstSell-price)
		firstSell = max(firstSell, firstBuy+price)
		firstBuy = max(firstBuy, -price)
	}

	return secondSell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
