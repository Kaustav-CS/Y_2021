/******************************************************************************
Link:   https://leetcode.com/problems/valid-number/
65. Valid Number

Your input
"0"

Output
true

Expected
true
DATE:    2021, October 25
০৮ কার্তিক,১৪২৮
*******************************************************************************/
func isNumber(s string) bool {
    idx := 0
    var numbers, eNotation, decimal bool
    
    if isSign(s[0]) { idx++ } // skip over any leading sign
    
    for ; idx < len(s); idx++ {
        if isDigit(s[idx]) {
            numbers = true
        } else {
            switch s[idx] {
            case 'e', 'E':
                if eNotation || !numbers {  return false } // there can only be one e notation and it must be preceeded by at least one number
                eNotation = true
                numbers = false // numbers bool is reset to false so we can check for digits following the e notation
                
                if idx < len(s) - 1 && isSign(s[idx+1]) { idx++ } // check for sign following e notation and skip over if so
            case '.':
                if eNotation || decimal { return false } // there can be only one decimal max, and it cannot come after any e notation
                decimal = true
            default:
                return false 
            }
        }
    }
    
    return numbers
}

// isDigit takes a byte and returns true if it is in the numerical range on the ASCII table, otherwise false
func isDigit(char byte) bool {
    return char >= '0' && char <= '9'
}

// isSign takes a byte and returns true if it corresponds to either '+' or '-', otherwise false
func isSign(char byte) bool {
    return char == '-' || char == '+'
}
