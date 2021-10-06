/******************************************************************************
Link:   https://leetcode.com/problems/implement-strstr/
28. Implement strStr()
Your input
"hello"
"ll"
Output
2
Expected
2

DATE:    2021 October 06
১৯ অশ্বিন, ১৪২৮

*******************************************************************************/
func strStr(haystack string, needle string) int {
    dict := make([]int, 128)
    a := haystack
    b := needle
    for i := 0; i < 26; i++ {
        dict[i] = -1
    }
    
    for i := 0; i < len(b); i++ {
        dict[b[i]] = i
    }
    
    i, j := 0, 0
    for i <= len(a) - len(b) {
        j = 0
        for j < len(b) {
            if a[i] == b[j] {
                i++
                j++
            } else {
                index := i + len(b) - j
                if index >= len(a) {
                    return -1
                }
                if dict[a[index]] == -1 {
                    i = index + 1
                } else {
                    i = index - dict[a[index]]
                }
                break
            }
        }
        if j == len(b) {
            return i - len(b)
        }
    }
    return -1
}
