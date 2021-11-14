/******************************************************************************
Link:   https://leetcode.com/problems/triangle/
120. Triangle

Your input
[[2],[3,4],[6,5,7],[4,1,8,3]]
Output
11
Expected
11

DATE:    2021, November 14
২৮ কার্তিক,১৪২৮

*******************************************************************************/
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    if n==0 {
        return 0
    }
    i:=0
    for i =1;i<n;i++ {
        prev := i-1
            length := len(triangle[prev])
            for j:= range triangle[i] {
                min := 2147483647
                if j == 0 {
                    min = triangle[prev][j]
                }else if j == length {
                    min = triangle[prev][j-1]
                }else {
                    if j-1 >=0 && triangle[prev][j-1] < triangle[prev][j] {
                        min = triangle[prev][j-1]
                    }else {
                        min = triangle[prev][j]
                    }
                }
                triangle[i][j] += min
            }
    }
    min := triangle[n-1][0]
    for i := 1; i<n; i++ {
        if triangle[n-1][i] < min {
            min = triangle[n-1][i]
        }
    }
    return min
}
