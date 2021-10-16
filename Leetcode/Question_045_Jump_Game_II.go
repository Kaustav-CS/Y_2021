/******************************************************************************
Link:   https://leetcode.com/problems/jump-game-ii/
45. Jump Game II

Your input
[2,3,1,1,4]

Output
2

Expected
2

DATE:    2021 October 16
২৯ অশ্বিন, ১৪২৮
*******************************************************************************/
func jump(nums []int) int {
    if len(nums) <= 1 {
		return 0
	}

	var maxIdx, i, step int
	for i < len(nums)-1 {
		nextIdx, tmpIdx := 0, 0
		for j := 0; j < nums[i]; j++ {
			tmpIdx = i + j + 1
			if tmpIdx < len(nums) && tmpIdx+nums[tmpIdx] > maxIdx {
				maxIdx = tmpIdx + nums[tmpIdx]
				nextIdx = tmpIdx
			}
			if tmpIdx == len(nums)-1 {
				nextIdx = tmpIdx
			}
		}

		i = nextIdx
		step++
	}

	return step
}
