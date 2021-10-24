/******************************************************************************
Link:   https://leetcode.com/problems/unique-paths-ii/
63. Unique Paths II

Your input
[[0,0,0],[0,1,0],[0,0,0]]

Output
2

Expected
2

DATE:    2021, October 24
০৭ কার্তিক,১৪২৮
*******************************************************************************/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    r := len(obstacleGrid)
    c := len(obstacleGrid[0])
    
    dp := make([][]int, r)
    for i := 0; i < r; i++ {
        dp[i] = make([]int, c)
    }
    
    for i := 0; i < r; i++ {
        if obstacleGrid[i][0] == 1 {
            break
        }
        dp[i][0] = 1
    }
    
    for i := 0; i < c; i++ {
        if obstacleGrid[0][i] == 1 {
            break
        }
        dp[0][i] = 1
    }
    
    for i := 1; i < r; i++ {
        for j := 1; j < c; j++ {
            if obstacleGrid[i][j] == 1 {
                continue
            }
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
        }
    }
    
    return dp[r - 1][c - 1]
}
