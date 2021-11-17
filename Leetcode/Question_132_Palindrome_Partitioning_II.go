/******************************************************************************
Link:   https://leetcode.com/problems/palindrome-partitioning-ii/
132. Palindrome Partitioning II

Your input
"aab"

Output
1

Expected
1

DATE:    2021, November 17
০১ অগ্রহায়ণ,১৪২৮
*******************************************************************************/
func minCut(s string) int {
    if len(s) == 0 {
        return 0
    }
    dict := make(map[int]int)
    return helper(s, dict)
}

func helper(s string, dict map[int]int) int {
    if len(s) == 0 {
        return -1
    }
    
    if v, ok := dict[len(s)]; ok {
        return v
    }
    
    result := math.MaxInt64
    for i := 1; i <= len(s); i++ {
        if isValid(s[:i]) {
            min := 1 + helper(s[i:], dict) 
            if min < result {
                result = min
            }
        }
    }   
    
    dict[len(s)] = result
    return result
}

func isValid(s string) bool {
    i, j := 0, len(s) - 1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

