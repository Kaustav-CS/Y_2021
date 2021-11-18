/******************************************************************************
Link:   https://leetcode.com/problems/single-number/
136. Single Number

Your input
[2,2,1]

Output
1

Expected
1

DATE:    2021, November 18
০২ অগ্রহায়ণ,১৪২৮

*******************************************************************************/
func singleNumber(nums []int) int {
    result := 0
    for _,v := range nums{
        result = result ^ v
    }
    return result
}
