/******************************************************************************
Link:   https://leetcode.com/problems/spiral-matrix-ii/
59. Spiral Matrix II

Your input
3

Output
[[1,2,3],[8,9,4],[7,6,5]]

Expected
[[1,2,3],[8,9,4],[7,6,5]]

DATE:    2021, October 23
০৬ কার্তিক,১৪২৮

*******************************************************************************/
func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0
	var num = 1
	for top <= bottom && left <= right {
		// traverse right
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// traverse down
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// traverse left
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// traverse up
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}
