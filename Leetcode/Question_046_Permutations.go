/******************************************************************************
Link:   https://leetcode.com/problems/permutations/
46. Permutations

Your input
[1,2,3]

Output
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Expected
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

DATE:    2021 October 16
২৯ অশ্বিন, ১৪২৮
*******************************************************************************/
func permute(nums []int) [][]int {
    if len(nums) == 0 {
		return nil
	}

	ans := make([][]int, 0)
	backtrack(nums, nil, &ans)

	return ans
}

func backtrack(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, prev...))
		return
	}

	for i := 0; i < len(nums); i++ {
		backtrack(append(append([]int{}, nums[0:i]...), nums[i+1:]...),
			append(prev, nums[i]),
			ans)
	}
}
