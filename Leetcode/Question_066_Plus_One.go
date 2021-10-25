/******************************************************************************
Link:   https://leetcode.com/problems/plus-one/
66. Plus One

Your input
[1,2,3]

Output
[1,2,4]

Expected
[1,2,4]

DATE:    2021, October 25
০৮ কার্তিক,১৪২৮
*******************************************************************************/
func plusOne(digits []int) []int {
    var n int = len(digits)
    for i:= n-1; i>=0; i--{
        if digits[i] < 9 {
            digits[i]+=1
            return digits
        } else {
            digits[i] = 0
        }
    }
    var a = make([]int,n+1)
    a[0] = 1
    return a

}
