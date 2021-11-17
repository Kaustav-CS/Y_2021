/******************************************************************************
Link:   https://leetcode.com/problems/palindrome-partitioning/
131. Palindrome Partitioning

Your input
"aab"

Output
[["a","a","b"],["aa","b"]]

Expected
[["a","a","b"],["aa","b"]]

DATE:    2021, November 17
০১ অগ্রহায়ণ,১৪২৮
*******************************************************************************/
func partition(s string) [][]string {
    ans := [][]string{}
    DFS(&ans, []string{}, 0, s)
    return ans
}

func DFS(ans *[][]string, cur []string, offset int, s string) {
    if len(s) == offset {
        *ans = append(*ans, append([]string{}, cur...))
        return
    }
    
    for i := offset; i<len(s); i++ {
        subStr := s[offset:i+1]
        if isPalindrome(subStr) {
            cur = append(cur, subStr)
            DFS(ans, cur, i+1, s)
            cur = cur[:len(cur)-1]
        }
    }
}

func isPalindrome(s string) bool {
    l := len(s)
    for i, j := 0, l-1; i<j; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    } 
    return true
}

