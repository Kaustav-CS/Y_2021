/******************************************************************************
Link:   https://leetcode.com/problems/maximum-product-subarray/
152. Maximum Product Subarray

Your input
[2,3,-2,4]

Output
6

Expected
6

DATE:    2021, December 27
১১ পৌষ,১৪২৮
*******************************************************************************/
func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    lastMax, lastMin := 1, 1
    maxVal := nums[0]
    for i := range nums {
        if nums[i] < 0 {
            lastMax, lastMin = lastMin, lastMax
        }
        
        // track max and min till the current num seperately
        // at each num, we can choose to use it as a new start or not, depends on the value.
        lastMax = max(lastMax*nums[i], nums[i])
        lastMin = min(lastMin*nums[i], nums[i])
        if lastMax > maxVal {
            maxVal = lastMax
        }
    }
    return maxVal
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
