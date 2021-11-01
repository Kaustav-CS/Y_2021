/******************************************************************************
Link:   https://leetcode.com/problems/decode-ways/
91. Decode Ways

Your input
"12"

Output
2

Expected
2

DATE:    2021, November 01
১৫ কার্তিক,১৪২৮

*******************************************************************************/
func numDecodings(s string) int {
    if s[0] == '0' {
        return 0
    }
    n := len(s)
    prevPrev := 1
    prev := 1
    for i := 1; i < n; i++ {
        curr := 0
        if s[i] != '0' {
            curr = prev
        }
        twoDigit, _ := strconv.Atoi(s[i-1:i+1])
        if twoDigit >= 10 && twoDigit <= 26 {
            curr += prevPrev
        }
        prevPrev = prev
        prev = curr
    }
    return prev
}

