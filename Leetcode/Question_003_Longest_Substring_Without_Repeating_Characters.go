/******************************************************************************
Link:   https://leetcode.com/problems/add-two-numbers/
Question_003_Longest_Substring_Without_Repeating_Characters

Given a string 's', 
find the length of the 'longest substring' without repeating characters.


Your input
"abcabcbb"

Output
3

Expected
3

*******************************************************************************/

func lengthOfLongestSubstring(s string) int {
    runes := []rune(s)
    lastIdx := make(map[rune]int, 0)
    maxLength := 0
    for i, j := 0, 0; j < len(runes); j += 1 {
        if last, ok := lastIdx[runes[j]]; ok && last >= i{
            i = last + 1            
        }
        if j - i + 1 > maxLength {
            maxLength = j - i + 1
        }                 
        lastIdx[runes[j]] = j
    }
    return maxLength
    
}
