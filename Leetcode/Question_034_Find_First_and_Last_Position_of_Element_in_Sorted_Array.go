func searchRange(nums []int, target int) []int {
    result := []int{-1, -1}
	i1, j1 := 0, len(nums)
	for i1 < j1 {
		h := int(uint(i1+j1) >> 1) // avoid overflow when computing h
		if nums[h] == target {
			if h == 0  {
				result[0] = h
				break
			} else if nums[h-1] != target {
				result[0] = h
				break
			}
		}
		if nums[h] < target {
			i1 = h + 1
		} else {
			j1 = h
		}
	}

	i2, j2 := 0, len(nums)
	for i2 < j2 {
		h := int(uint(i2+j2) >> 1)
		if nums[h] == target {
			if h == len(nums) - 1  {
				result[1] = h
				break
			} else if nums[h+1] != target {
				result[1] = h
				break
			}
		}
		if nums[h] <= target {
			i2 = h + 1
		} else {
			j2 = h
		}
	}

	return result
}
