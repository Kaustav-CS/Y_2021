/******************************************************************************
Link:   https://leetcode.com/problems/single-number-ii/
137. Single Number II

Your input
[2,2,3,2]

Output
3

Expected
3

DATE:    2021, November 18
০২ অগ্রহায়ণ,১৪২৮

*******************************************************************************/
func singleNumber(nums []int) int {
    a, b := 0, 0
    for _, num := range nums {
        a = a ^ num & ^b
        b = b ^ num & ^a
    }
    return a | b
}

