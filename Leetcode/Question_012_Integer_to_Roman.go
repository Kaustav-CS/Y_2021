/******************************************************************************
Link:   https://leetcode.com/problems/integer-to-roman/
12. Integer to Roman

 Question_012_Integer_to_Roman.go

Your input
3

Output
"III"

Expected
"III"

DATE:    2021 September 28
১১ অশ্বিন, ১৪২৮

*******************************************************************************/
func intToRoman(num int) string {
    values := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    symbols := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    
    builder := strings.Builder{}
    for i, v := range values {
        for num >= v {
            num -= v
            builder.WriteString(symbols[i])
        }
    }
    
    return builder.String()
}
