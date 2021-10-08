/******************************************************************************
Link:   https://leetcode.com/problems/next-permutation/
31. Next Permutation

Your input
[1,2,3]

Output
[1,3,2]

Expected
[1,3,2]


DATE:    08/10/2021
২১ অশ্বিন, ১৪২৮

*******************************************************************************/
func nextPermutation(nums []int)  {
    length := len(nums)
	// test if nums[i] < nums[i-1] true
	for i := length-2; i >= 0; i-- {
		if nums[i] < nums[i + 1] {
			// i < swapIndex < length && nums[swapIndex] > nums[i] && nearest
			swapIndex := i + 1
			for j := i+2; j < length; j++ {
				if nums[j] > nums[i] {
					swapIndex = j
				} else {
					break
				}
			}
			// swap nums[i], nums[swapIndex]
			nums[i], nums[swapIndex] = nums[swapIndex], nums[i]
			// reverse nums[i+1:]
			for j := i+1; j < i+1+(length-i-1)/2; j++ {
				nums[j], nums[length-j+i] = nums[length-j+i], nums[j]
			}
			return
		}
	}
	// just reverse it
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
	}
}
