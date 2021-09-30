/******************************************************************************
Link:   https://leetcode.com/problems/3sum/
15. 3Sum


Your input
[-1,0,1,2,-1,-4]

Output
[[-1,-1,2],[-1,0,1]]

Expected
[[-1,-1,2],[-1,0,1]]


DATE:    2021 October 01
১৪ অশ্বিন, ১৪২৮
*******************************************************************************/
func threeSum(nums []int) [][]int {
    ans := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {continue}
        j, k := i + 1, len(nums) - 1
        for j < k {
            if j > i + 1 && nums[j] == nums[j-1] {
                j++
                continue
            }
            if k < len(nums) - 1 && nums[k] == nums[k+1] {
                k--
                continue
            }
            sum := nums[i] + nums[j] + nums[k]
            if sum == 0 {
                ans = append(ans, []int{nums[i], nums[j], nums[k]})
                j++
                k--
            } else if sum < 0 {
                j++
            } else {
                k--
            }
        }
    }    
    return ans
}
