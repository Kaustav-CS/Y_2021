/******************************************************************************
Link:   https://leetcode.com/problems/generate-parentheses/
22. Generate Parentheses

Your input
3

Output
["((()))","(()())","(())()","()(())","()()()"]

Expected
["((()))","(()())","(())()","()(())","()()()"]

DATE:    2021 October 04
১৭ অশ্বিন, ১৪২৮
*******************************************************************************/
func generateParenthesis(n int) []string {
    res := make([]string, 0)
    backtrack(n, n, &res, "")
    return res
}

func backtrack(left, right int, res *[]string, cur string) {
    if left == 0 && right == 0 {
        *res = append(*res, cur)
        return
    }
    
    if right < left {
        return
    }
    
    if left > 0 {
        backtrack(left-1, right, res, cur+"(")
    }

    if right > 0 {
        backtrack(left, right-1, res, cur+")")
    }
}   
