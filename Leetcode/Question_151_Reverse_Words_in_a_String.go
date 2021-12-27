/******************************************************************************
Link:   https://leetcode.com/problems/reverse-words-in-a-string/
151. Reverse Words in a String

Your input
"the sky is blue"

Output
"blue is sky the"

Expected
"blue is sky the"

DATE:    2021, December 27
১১ পৌষ,১৪২৮
*******************************************************************************/
func reverseWords(s string) string {
    i, n, res := 0, len(s), ""
    for i < n {
        for i < n && s[i] == ' ' { i++ }
        if i == n { break }
        j := i
        for j < n && s[j] != ' ' { j++ }
        if len(res) == 0 {
            res = s[i : j]
        } else {
            res = s[i : j] + " " + res
        }
        i = j + 1
    }
    return res
}
