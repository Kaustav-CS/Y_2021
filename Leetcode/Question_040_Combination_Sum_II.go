/******************************************************************************
Link:   https://leetcode.com/problems/combination-sum-ii/
40. Combination Sum II

Your input
[10,1,2,7,6,1,5]
8

Output
[[1,1,6],[1,2,5],[1,7],[2,6]]

Expected
[[1,1,6],[1,2,5],[1,7],[2,6]]


DATE:    2021 October 14
২৭ অশ্বিন, ১৪২৮


*******************************************************************************/
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    result := make([][]int, 0)
    combinationSum2Helper(candidates, target, []int{}, &result)
    return result
}

func combinationSum2Helper(candidates []int, target int, current []int, result *[][]int) {
	if target <= 0 || len(candidates) == 0 {
		return
	}
	for i, v := range candidates {
        if v >= target {
            if v == target {
                *result = append(*result, append(current, v))
            }
            break
        } else if i != 0 && (i == len(candidates)-1 || candidates[i] != candidates[i+1]) {
            combinationSum2Helper(candidates[:i], target-v, append(current, v), result)
        }
	}
	return
}
