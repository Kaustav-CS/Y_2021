/******************************************************************************
Link:   https://leetcode.com/problems/first-missing-positive/
41. First Missing Positive

Your input
[1,2,0]

Output
3

Expected
3


DATE:    2021 October 15
২৮ অশ্বিন, ১৪২৮



*******************************************************************************/
func firstMissingPositive(nums []int) int {
    sz := len(nums)
    
    for i := 0; i < sz; i++ {        
        // ensure no negative
        if nums[i] <= 0 {
            nums[i] = sz + 1
        }
    }
    
    for i := 0; i < sz; i++ {
        n := int(math.Abs(float64(nums[i])))
        
        if n > sz {
            continue
        }
        
        if nums[n - 1] > 0 {
            nums[n - 1] *= -1
        }
    }
    
    for i := 1; i <= sz; i++ {
        if nums[i - 1] > 0 {
            return i
        }
    }    
    
    
    return sz + 1
}
