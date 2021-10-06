/******************************************************************************
Link:   https://leetcode.com/problems/remove-element/
27. Remove Element

Your input
[3,2,2,3]
3

Output
[2,2]

Expected
[2,2]

DATE:    2021 October 06
১৯ অশ্বিন, ১৪২৮

*******************************************************************************/
func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    index := 0
    for i := 0; i < len(nums); i++  {
        if nums[i] != val {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}
