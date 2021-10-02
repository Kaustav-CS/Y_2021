/******************************************************************************
Link:   https://leetcode.com/problems/letter-combinations-of-a-phone-number/
17. Letter Combinations of a Phone Number


Given a string containing digits from 2-9 inclusive, return all possible 
letter combinations that the number could represent. 
Return the answer in any order.

Your input
"23"

Output
["ad","ae","af","bd","be","bf","cd","ce","cf"]

Expected
["ad","ae","af","bd","be","bf","cd","ce","cf"]

DATE:    2021 October 02
১৫ অশ্বিন, ১৪২৮
*******************************************************************************/
func letterCombinations(digits string) []string {
    ans := []string{}
    if len(digits) == 0 {return ans}
    helper(digits, 0, "", &ans)
    return ans
}

var M = map[byte]string {
    '1': "",
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
    '0': " ",
}

func helper(digits string, pos int, s string, strs *[]string) {
    if pos == len(digits) {
        *strs = append(*strs, s)
        return
    }
    
    for _, c := range M[digits[pos]] {
        helper(digits, pos + 1, s + string(c), strs)
    }
    return
}
