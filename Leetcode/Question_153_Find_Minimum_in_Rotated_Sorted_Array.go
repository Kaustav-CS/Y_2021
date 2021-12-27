/******************************************************************************
Link:   https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
153. Find Minimum in Rotated Sorted Array

Your input
[3,4,5,1,2]

Output
1

Expected
1

DATE:    2021, December 27
১১ পৌষ,১৪২৮
*******************************************************************************/
func findMin(nums []int) int {
    start,end := 0,len(nums)-1
	var mid int
	for start < end{
		if nums[start] < nums[end]{
			return nums[start]
		}
		mid = (start + end) / 2
		if nums[mid] > nums[end]{
			start = mid + 1
		}else{
			end = mid
		}
	}
	return nums[start]
}
