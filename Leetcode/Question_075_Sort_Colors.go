/******************************************************************************
Link:   https://leetcode.com/problems/sort-colors/
75. Sort Colors

Your input
[2,0,2,1,1,0]

Output
[0,0,1,1,2,2]

Expected
[0,0,1,1,2,2]

DATE:    2021, October 28
১১ কার্তিক,১৪২৮

*******************************************************************************/
func sortColors(nums []int)  {
    zeros, ones, twos := 0, 0, 0
    
    for i := 0; i < len(nums); i++ {
        x := nums[i]

        if x == 0 {
            nums[zeros], nums[i] = nums[i], nums[zeros]
            zeros++
        } else if x == 1 {
            nums[zeros + ones], nums[i] = nums[i], nums[zeros + ones]
            ones++
        } else if x == 2 {
            nums[zeros + ones + twos], nums[i] = nums[i], nums[zeros + ones + twos]
            twos++
        }
        
        for j := i; j > 0 && nums[j] < nums[j - 1]; j-- {
            nums[j], nums[j-1] = nums[j-1], nums[j]
        }
    }
}
