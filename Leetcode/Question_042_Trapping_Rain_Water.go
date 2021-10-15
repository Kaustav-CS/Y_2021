/******************************************************************************
Link:   https://leetcode.com/problems/trapping-rain-water/
42. Trapping Rain Water

Your input
[0,1,0,2,1,0,1,3,2,1,2,1]

Output
6

Expected
6


DATE:    2021 October 15
২৮ অশ্বিন, ১৪২৮



*******************************************************************************/
func trap(height []int) int {
    left, right, sum, maxLeft, maxRight := 0, len(height)-1, 0, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				sum += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				sum += maxRight - height[right]
			}
			right--
		}
	}
	return sum
}
