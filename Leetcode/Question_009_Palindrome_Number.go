/******************************************************************************
''
Link:   https://leetcode.com/problems/palindrome-number/


9. Palindrome Number

Your input
121
Output
true
Expected
true

DATE:    2021 September 27
১০ অশ্বিন, ১৪২৮


*******************************************************************************/
func isPalindrome(x int) bool {
    if x<0 {
        return false
    }
    
    if x==0{
        return true
    }
    
    var reverse, origin int
    origin = x
    
    for x!=0{
        val:= x%10
        reverse = reverse*10+val
        x = x/10
    }
    
    return origin==reverse
}
