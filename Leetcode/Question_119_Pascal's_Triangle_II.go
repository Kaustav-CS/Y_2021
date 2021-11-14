/******************************************************************************
Link:   https://leetcode.com/problems/pascals-triangle-ii/
119. Pascal's Triangle II

Your input
3
Output
[1,3,3,1]
Expected
[1,3,3,1]

DATE:    2021, November 14
২৮ কার্তিক,১৪২৮

*******************************************************************************/
func getRow(rowIndex int) []int {
    return generate(rowIndex + 1)[rowIndex]
}
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := make([][]int, numRows)
	for j := 0; j <= numRows-1; j++ {

		result[j] = make([]int, j+1)
		if numRows == 1 {
			result[0][0] = 1
			return result
		}
		for i := 0; i <= j; i++ {
			if i == 0 || i == j {
				result[j][i] = 1
			} else {
				result[j][i] = result[j-1][i-1] + result[j-1][i]
			}
		}
	}

	return result
}
