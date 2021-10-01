/******************************************************************************
Link:   https://leetcode.com/problems/3sum-closest/
16. 3Sum Closest
Question_016_3Sum_Closest.go

Your input
[-1,2,1,-4]
1

Output
2

Expected
2

DATE:    2021 October 01
১৪ অশ্বিন, ১৪২৮
*******************************************************************************/
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    ans := nums[0] + nums[1] + nums[2]
    n := len(nums)    
    for i := 0; i < n; i++ {
        j, k := i + 1, n - 1
        for j < k {
            v := nums[i] + nums[j] + nums[k]
            if abs(target - v) < abs(target - ans) {
                ans = v
            }
            if v <= target {
                j++
            } else {
                k--
            }
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
