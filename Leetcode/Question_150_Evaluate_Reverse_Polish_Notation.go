/******************************************************************************
Link:   https://leetcode.com/problems/evaluate-reverse-polish-notation/
150. Evaluate Reverse Polish Notation

Your input
["2","1","+","3","*"]

Output
9

Expected
9

DATE:    2021, December 27
১১ পৌষ,১৪২৮
*******************************************************************************/
func evalRPN(tokens []string) int {
    stk := []int{}
    
    none := struct{}{}
    
    operators := map[string]struct{} {
        "+": none,
        "-": none,
        "*": none,
        "/": none,
    }
    
    for _, val := range tokens {
        if _, ok := operators[val]; ok {
            // found
            // pop two 
            last := len(stk) - 1
            first := stk[last]
            second := stk[last - 1]
            
            // do opeation
            v := action(second, first, val) 
            stk = stk[:last - 1]
            stk = append(stk, v)
            
        } else {
            // push on stack
            v, _ := strconv.ParseInt(val, 10, 32)
            stk = append(stk, int(v))
        }
    }
    
    
    return stk[0]
}

func action(i, j int, op string) int {
    switch op {
        case "+":
            return i + j
    case "-":
        return i - j
        case "*":
        return i * j
    case "/": 
        return i / j
    }    
    
    return 0
}

