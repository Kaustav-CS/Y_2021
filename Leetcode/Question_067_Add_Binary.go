/******************************************************************************
Link:   https://leetcode.com/problems/add-binary/
67. Add Binary

Your input
"11"
"1"

Output
"100"

Expected
"100"

DATE:    2021, October 26
০৯ কার্তিক,১৪২৮

*******************************************************************************/
func addBinary(a string, b string) string {
    s := 0
    carry := 0
    res := ""
    la := len(a) - 1
    lb := len(b) - 1
    for la >= 0 || lb >= 0 || carry != 0{
        s = carry
        if la >= 0 {
            s += int(a[la] - '0')
            la -- 
        }
        if lb >= 0 {
            s += int(b[lb] - '0')
            lb --
        }
        carry = s / 2
        res = string(s % 2 + '0') + res
    }
    return res
}
