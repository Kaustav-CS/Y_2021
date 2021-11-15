/******************************************************************************
Link:   https://leetcode.com/problems/valid-palindrome/
125. Valid Palindrome

Your input
"A man, a plan, a canal: Panama"

Output
true

Expected
true

{
"A man, a plan, a canal: Panama"
"race a car"
}

DATE:    2021, November 15
২৯ কার্তিক,১৪২৮

*******************************************************************************/
func isPalindrome(s string) bool {
    for i,j := 0, len(s) - 1; i < len(s); {
        if !isAlphanumeric(s[i]) {
            i++ 
            continue
        }   

        if !isAlphanumeric(s[j]) {
            j-- 
            continue
        }   

        if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
            return false
        }   

        i++ 
        j-- 
    }   

    return true
}

func isAlphanumeric(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
