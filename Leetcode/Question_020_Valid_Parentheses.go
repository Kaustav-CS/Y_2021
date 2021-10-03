/******************************************************************************
Link:   https://leetcode.com/problems/valid-parentheses/
20. Valid Parentheses

Your input
"()"

Output
true

Expected
true

DATE:    2021 October 03
১৬ অশ্বিন, ১৪২৮
*******************************************************************************/
func isValid(s string) bool {
    seen := map[rune]rune{
        '(' : ')',
        '[' : ']',
        '{' : '}',
        ')' : '(',
        ']' : '[',
        '}' : '{',
    }
    
    stk := []rune{}
    
    for _, i := range s {
        if i == ')' || i == '}' || i == ']' {
            if len(stk) > 0 && seen[i] == stk[len(stk) - 1] {
                stk = stk[: len(stk) - 1]
                continue
            }
        }
        
        stk = append(stk, i)
    }
    
    return len(stk) == 0
}
