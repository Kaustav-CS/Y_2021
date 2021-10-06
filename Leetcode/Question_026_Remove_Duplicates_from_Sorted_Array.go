/******************************************************************************
Link:   https://leetcode.com/problems/remove-duplicates-from-sorted-array/
26. Remove Duplicates from Sorted Array

Your input
[1,1,2]

Output
[1,2]

Expected
[1,2]

DATE:    2021 October 06
১৯ অশ্বিন, ১৪২৮

*******************************************************************************/
func removeDuplicates(nums []int) int {
    nMap := make(map[int]int)
    k := 0
    for i := 0; i < len(nums); i++ {
        _, ok := nMap[nums[i]]
        if ok {
            continue
        }
        nMap[nums[i]] = k
        nums[k] = nums[i]
        k++
    }
    return(len(nums[:k]))
}
