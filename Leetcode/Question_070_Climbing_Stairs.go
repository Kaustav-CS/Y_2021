/******************************************************************************
Link:   https://leetcode.com/problems/climbing-stairs/
70. Climbing Stairs

Your input
2

Output
2

Expected
2

DATE:    2021, October 26
০৯ কার্তিক,১৪২৮

*******************************************************************************/
func climbStairs(n int) int {
    return fibonacci(n+1)
}

func fibonacci(n int) int {
    var x, y int = 0, 1
    for n > 0 {
        x, y = y, x+y
        n--
    }
    return x
}
