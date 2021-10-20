/******************************************************************************
Link:   https://leetcode.com/problems/jump-game/
55. Jump Game

Your input
[2,3,1,1,4]

Output
true

Expected
true


DATE:    2021, October 20
০৩ কার্তিক,১৪২৮
*******************************************************************************/
func canJump(nums []int) bool {
    length := len(nums)
	if length == 1 {
		return true
	} else if nums[0] == 0 && length > 1 {
		return false
	}

	maxJump := nums[0]
	for i := 1; i < length-1; i++ {
		next := nums[i] + i
		if next > maxJump {
			maxJump = next
		}

		if nums[i] == 0 && i >= maxJump {
			return false
		}
	}

	return true
}
