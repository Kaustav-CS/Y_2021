/******************************************************************************
Given two sorted arrays 'nums1' and 'nums2' of size 'm' and 'n' respectively, 
return 'the median' of the two sorted arrays.

The overall run time complexity should be 'O(log (m+n))'.

Link:   https://leetcode.com/problems/longest-palindromic-substring/

5. Longest Palindromic Substring

Your input
"babad"
Output
"bab"
Expected
"bab"
*******************************************************************************/
func longestPalindrome(s string) string {
    lo, maxLen := 0, 0
	if len(s) < 2 {
		return s
	}

	for i := range s {
		extendPal(s, i, i, &lo, &maxLen)
		extendPal(s, i, i+1, &lo, &maxLen)
	}

	return s[lo:lo+maxLen]
}
