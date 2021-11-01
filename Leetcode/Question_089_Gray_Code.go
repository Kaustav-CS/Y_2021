/******************************************************************************
Link:   https://leetcode.com/problems/gray-code/
89. Gray Code

Your input
2

Output
[0,1,3,2]

Expected
[0,1,3,2]

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
