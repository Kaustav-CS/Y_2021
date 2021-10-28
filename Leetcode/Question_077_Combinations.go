/******************************************************************************
Link:   https://leetcode.com/problems/combinations/
77. Combinations

Your input
4
2

Output
[[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]

Expected
[[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]

DATE:    2021, October 28
১১ কার্তিক,১৪২৮

*******************************************************************************/
func combine(n int, k int) [][]int {
    result := make([][]int, 0)
	helper(1, n, k, []int{}, &result)
	return result
}

func helper(pointer, n, k int, current []int,  result *[][]int) {
	if len(current) == k {
		*result = append(*result, append([]int{}, current...))
		return
	}
	for i := pointer; i <= n; i++ {
		helper(i+1, n, k, append(current, i), result)
	}
}
