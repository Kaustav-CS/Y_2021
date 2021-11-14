/******************************************************************************
Link:   https://leetcode.com/problems/pascals-triangle/
118. Pascal's Triangle

Your input
5
Output
[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Expected
[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

DATE:    2021, November 14
২৮ কার্তিক,১৪২৮

*******************************************************************************/
func generate(numRows int) [][]int {
    if numRows == 0 {
		return [][]int{}
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	if numRows == 1 {
		return res
	}
	for i := 1; i <numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 0; j < len(res[i-1])-1; j++ {
			tmp[j+1] = res[i-1][j] + res[i-1][j+1]
		}
		res[i] = tmp
	}
	return res
}
