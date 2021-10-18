/******************************************************************************
Link:   https://leetcode.com/problems/rotate-image/
48. Rotate Image

Your input
[[1,2,3],[4,5,6],[7,8,9]]

stdout
[[7 4 1] [8 5 2] [9 6 3]]

Output
[[7,4,1],[8,5,2],[9,6,3]]

Expected
[[7,4,1],[8,5,2],[9,6,3]]

DATE:    2021 October 18
০১ কার্তিক,১৪২৮
*******************************************************************************/
func rotate(matrix [][]int)  {
    c := len(matrix[0])
	r := len(matrix)
	for i := 0; i<c; i++ {
		for j := i; j< r  ; j++{
			tmp := matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
	for i := 0; i<c; i++ {
		for j := r/2; j< r; j++{
				tmp := matrix[i][r - j -1]
				matrix[i][r - j -1] = matrix[i][j]
				matrix[i][j] = tmp 
		}
	}
	fmt.Println(matrix)
}
