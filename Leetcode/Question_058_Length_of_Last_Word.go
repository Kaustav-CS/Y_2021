/******************************************************************************
Link:   https://leetcode.com/problems/length-of-last-word/
58. Length of Last Word

Your input
"Hello World"

Output
5

Expected
5


DATE:    2021, October 22
০৫ কার্তিক,১৪২৮
*******************************************************************************/
import "strings"

func lengthOfLastWord(s string) int {
    length := 0
    s = strings.TrimSpace(s) // testcase like "XXX ", WTF
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == ' ' {
            break
        } else {
            length++
        }
    }
    return length
}
