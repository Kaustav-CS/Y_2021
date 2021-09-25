/******************************************************************************
Given two sorted arrays 'nums1' and 'nums2' of size 'm' and 'n' respectively, 
return 'the median' of the two sorted arrays.

The overall run time complexity should be 'O(log (m+n))'.

Link:   https://leetcode.com/problems/median-of-two-sorted-arrays/

4. Median of Two Sorted Arrays
*******************************************************************************/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
 nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	l := len(nums1)

	var median float64
	if l%2 == 1 {
		median = float64(nums1[l/2.0])
	} else {
		var n1 float64
		var n2 float64
		n1 = float64(nums1[l/2])
		n2 = float64(nums1[(l/2)-1.0])

		median = (n1 + n2) / 2.0
	}
	return float64(median)   
}
