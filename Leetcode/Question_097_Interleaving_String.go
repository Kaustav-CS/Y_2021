/******************************************************************************
Link:   https://leetcode.com/problems/interleaving-string/
97. Interleaving String

Your input
"aabcc"
"dbbca"
"aadbbcbcac"

Output
true

Expected
true

DATE:    2021, November 03
১৭ কার্তিক,১৪২৮

*******************************************************************************/
func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1)+len(s2)!=len(s3){
        return false
    }
    dp:=make([][]bool,len(s1)+1)
    for i:=0;i<len(s1)+1;i++{
        dp[i]=make([]bool,len(s2)+1)
    }
    dp[0][0]=true
    for i:=1;i<len(s1)+1;i++{
        if s3[i-1]==s1[i-1]{
            dp[i][0]=dp[i-1][0]
        }
    }
    for i:=1;i<len(s2)+1;i++{
        if s3[i-1]==s2[i-1]{
            dp[0][i]=dp[0][i-1]
        }
    }
    for i:=1;i<len(s1)+1;i++{
        for j:=1;j<len(s2)+1;j++{
            dp[i][j]=(s1[i-1]==s3[i+j-1]&&dp[i-1][j])||(s2[j-1]==s3[i+j-1]&&dp[i][j-1])
        }
    }
    return dp[len(s1)][len(s2)]
}

