/******************************************************************************
Link:   https://leetcode.com/problems/maximum-subarray/
53. Maximum Subarray

Your input
[-2,1,-3,4,-1,2,1,-5,4]

Output
6

Expected
6

DATE:    2021, October 20
০৩ কার্তিক,১৪২৮
*******************************************************************************/
func maxSubArray(nums []int) int {
    sum := nums[0]
    if sum < 0 {
        sum = 0
    }
    max := nums[0]
    for _,v := range nums[1:] {
        sum = sum + v
        if sum > max {
            max = sum
        }
        if sum < 0 {
            sum = 0
        } 
    }
    return max
}
