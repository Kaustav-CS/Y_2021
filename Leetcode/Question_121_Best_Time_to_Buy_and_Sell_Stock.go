/******************************************************************************
Link:   https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
121. Best Time to Buy and Sell Stock

Your input
[7,1,5,3,6,4]

Output
5

Expected
5

DATE:    2021, November 15
২৯ কার্তিক,১৪২৮

*******************************************************************************/
func maxProfit(prices []int) int {
    maxSeen := 0
    maxDiffSeen := 0
    
    for i := len(prices)-1 ; i >= 0; i-- {
        if prices[i] > maxSeen {
            maxSeen = prices[i]
        }
        if maxSeen-prices[i] > maxDiffSeen {
            maxDiffSeen = maxSeen-prices[i]
        }
    }
    
    return maxDiffSeen
}
