/******************************************************************************
Link:   https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
80. Remove Duplicates from Sorted Array II

Your input
[1,1,1,2,2,3]

Output
[1,1,2,2,3]

Expected
[1,1,2,2,3]


DATE:    2021, October 29
১২ কার্তিক,১৪২৮

*******************************************************************************/
func removeDuplicates(nums []int) int {
    cur := 10001 // cur must be different than nums[0]
    l := 0
    allow := true
    for _, n := range nums {
        if n != cur || allow {
            allow = n != cur
            cur, nums[l] = n, n
            l++
        }
    }
    return l
}
