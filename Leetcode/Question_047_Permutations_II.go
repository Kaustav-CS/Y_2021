/******************************************************************************
Link:   https://leetcode.com/problems/permutations-ii/
47. Permutations II

Your input
[1,1,2]

Output
[[1,1,2],[1,2,1],[2,1,1]]

Expected
[[1,1,2],[1,2,1],[2,1,1]]

DATE:    2021 October 16
২৯ অশ্বিন, ১৪২৮
*******************************************************************************/
func permuteUnique(nums []int) [][]int {
    if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	ans := make([][]int, 0)
	backtrackUnique(nums, nil, &ans)

	return ans
}

func backtrackUnique(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, prev...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		backtrackUnique(append(append([]int{}, nums[0:i]...), nums[i+1:]...),
			append(prev, nums[i]),
			ans)
	}
}
