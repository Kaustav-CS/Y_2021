/******************************************************************************
Link:   https://leetcode.com/problems/merge-sorted-array/
88. Merge Sorted Array

Your input
[1,2,3,0,0,0]
3
[2,5,6]
3

Output
[1,2,2,3,5,6]

Expected
[1,2,2,3,5,6]

DATE:    2021, November 01
১৫ কার্তিক,১৪২৮

*******************************************************************************/
func merge(nums1 []int, m int, nums2 []int, n int)  {
    for n > 0 {
		if m == 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}
