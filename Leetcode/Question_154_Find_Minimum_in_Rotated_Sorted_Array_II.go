/******************************************************************************
Link:   https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
154. Find Minimum in Rotated Sorted Array II

Your input
[1,3,5]
Output
1
Expected
1

DATE:    2021, December 30
১৪ পৌষ,১৪২৮
*******************************************************************************/
func findMin(nums []int) int {
    left, right := 0, len(nums)-1
	for left < right {
		mid := (left+right)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

