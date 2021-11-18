/******************************************************************************
Link:   https://leetcode.com/problems/candy/
135. Candy

Your input
[1,0,2]

Output
5

Expected
5

DATE:    2021, November 18
০২ অগ্রহায়ণ,১৪২৮

*******************************************************************************/
func candy(ratings []int) int {
    candies := len(ratings)
    left, right := 0, 0
    
    for i := 1; i < len(ratings); i++ {
        if ratings[i-1] < ratings[i]{
            left++
            candies += left
            continue
        }
        if ratings[i-1] == ratings[i]{
            left=0
            continue
        }
        candies += right
        right++
        if i+1 == len(ratings) || ratings[i] <= ratings[i+1] {
            candies += max(right-left,0)
            left,right = 0,0
        }
    }
    return candies
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

