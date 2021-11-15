/******************************************************************************
Link:   https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
122. Best Time to Buy and Sell Stock II
{122 of 2076}
Your input
[7,1,5,3,6,4]

Output
7

Expected
7

DATE:    2021, November 15
২৯ কার্তিক,১৪২৮

*******************************************************************************/
func maxProfit(prices []int) int {
    /* if len <= 1, there is no transaction
       calculate all the raising phase difference
    */
    if len(prices) <= 1 {
        return 0
    }
    
    i, j, profit := 0, 1, 0
    for i < len(prices) - 1 && j < len(prices) {
        if prices[j] > prices[i] {
            profit = profit + prices[j] - prices[i]
            i = j
            j = j + 1
        } else {
            i = j
            j = i + 1
        }
    }
    
    return profit
}
