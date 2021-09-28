/******************************************************************************
Link:   https://leetcode.com/problems/container-with-most-water/
11. Container With Most Water

Your input
[1,8,6,2,5,4,8,3,7]

Output
49

Expected
49

DATE:    2021 September 28
১১ অশ্বিন, ১৪২৮


*******************************************************************************/
func maxArea(height []int) int {
 i, j := 0, len(height)-1
    res := 0
    for i < j {
        res = max(res, (j-i)*min(height[i], height[j]))
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    return res   
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
