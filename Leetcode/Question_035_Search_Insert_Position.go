/******************************************************************************
Link:   https://leetcode.com/problems/search-insert-position/
35. Search Insert Position

Your input
[1,3,5,6]
5
Output
2
Expected
2


DATE:    2021 October 10
২৩ অশ্বিন, ১৪২৮
*******************************************************************************/
func searchInsert(nums []int, target int) int {
    i := sort.Search(len(nums), func(i int) bool {
        return nums[i] >= target
    })
    return i
}
