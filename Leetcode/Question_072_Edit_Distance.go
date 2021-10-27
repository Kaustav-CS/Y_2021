/******************************************************************************
Link:   https://leetcode.com/problems/edit-distance/
72. Edit Distance

Your input
"horse"
"ros"

Output
3

Expected
3

DATE:    2021, October 27
১০ কার্তিক,১৪২৮
*******************************************************************************/
func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(word1, word2 int) int {
	if word1 < word2 {
		return word1
	}
	return word2
}
