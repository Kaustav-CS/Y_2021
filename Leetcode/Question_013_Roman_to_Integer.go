
/******************************************************************************
Link:   https://leetcode.com/problems/roman-to-integer/
13. Roman to Integer

Your input
"III"

Output
3

Expected
3

DATE:    2021 September 29
১২ অশ্বিন, ১৪২৮

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
