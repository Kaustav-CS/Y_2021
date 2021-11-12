/******************************************************************************
Link:   https://leetcode.com/problems/distinct-subsequences/
115. Distinct Subsequences

Your input
"rabbbit"
"rabbit"

Output
3

Expected
3

DATE:    2021, November 12
২৬ কার্তিক,১৪২৮
*******************************************************************************/
func numDistinct(s string, t string) int {
    d := make([][]int, len(s)+1)
    
    for i := 0; i <= len(s); i++ {
        d[i] = make([]int, len(t)+1)
        d[i][len(t)] = 1
    }
    
    for i := len(s)-1; i >= 0; i-- {
        for j := 0; j < len(t); j++ {
            d[i][j] = d[i+1][j]
            
            if s[i] == t[j] {
                d[i][j] += d[i+1][j+1]
            }
        }
    }
    
    return d[0][0]
}
