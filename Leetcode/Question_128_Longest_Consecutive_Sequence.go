/******************************************************************************
Link:   https://leetcode.com/problems/longest-consecutive-sequence/
128. Longest Consecutive Sequence

Your input
[100,4,200,1,3,2]

Output
4

Expected
4

DATE:    2021, November 16
৩০ কার্তিক,১৪২৮

*******************************************************************************/
func longestConsecutive(nums []int) int {
    if len(nums) <= 1 {
		return len(nums)
	}
	// Build the map, and in this way we also remove duplicates, takes O(n) to create the map
	mp := map[int]bool{}
	for _, n := range nums {
		mp[n] = true
	}

	maxL := 1

	for k := range mp {
		curL := 1
		// flags to indicate if we should continue to expand towards left or right
		toLeft, toRight := true, true
		for i := 1; mp[k-i] || mp[k+i]; i++ {
			// when the expanded number doesn't exist in the map, we stop expanding towards that direction
			toLeft = toLeft && mp[k-i]
			toRight = toRight && mp[k+i]
			if toLeft {
				curL++
				delete(mp, k-i)
			}
			if toRight {
				curL++
				delete(mp, k+i)
			}
		}
		if curL > maxL {
			maxL = curL
		}
		// always remove numbers that have been processed from the map
		delete(mp, k)
	}

	return maxL
}

