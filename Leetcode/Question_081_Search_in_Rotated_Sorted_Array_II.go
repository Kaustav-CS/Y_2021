/******************************************************************************
Link:   https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
81. Search in Rotated Sorted Array II

Your input
[2,5,6,0,0,1,2]
0

Output
true

Expected
true

DATE:    2021, October 29
১২ কার্তিক,১৪২৮

*******************************************************************************/
func search(nums []int, target int) bool {
    Lnums := len(nums)
	if Lnums == 0{
		return false
	}

	if Lnums == 1{
		if nums[0] == target {
			return true
		}else{
			return false
		}
	}

	left,right := 0,Lnums-1

	for left<right{
		// be attention, this may cause left equals to right,when all the nummbers are the same  
		for left<right && nums[left] == nums[left+1]{
			left++
		}
		for left<right && nums[right] == nums[right-1]{
			right--
		}

		if left == right{
			if nums[left] == target{
				return true
			}else{
				return false
			}
		}

		mid := (left + right) / 2
		if nums[mid] == target{
			return true
		}

		if nums[left] <= nums[mid]{
			if target >= nums[left] && target < nums[mid]{
				right = mid -1
			}else{
				left = mid + 1
			}
		}else{
			if target > nums[mid] && target <= nums[right]{
				left = mid + 1
			}else{
				right = mid - 1
			}
		}
	}
	
	if nums[left] == target{
		return true
	}else{
		return false
	}
}
