/******************************************************************************
Link:   https://leetcode.com/problems/n-queens-ii/
52. N-Queens II

Your input
4
Output
2
Expected
2

DATE:    2021, October 19
০২ কার্তিক,১৪২৮

*******************************************************************************/
func totalNQueens(n int) int {
    var count int
	columns := make([]bool, n)  // columns   |
	d1 := make([]bool, 2*n)     // diagonals \
	d2 := make([]bool, 2*n)		// diagonals /
	backtracking(0, n, columns, d1, d2, &count)
	return count
}

func backtracking(currentRow, n int, columns, d1, d2 []bool ,count *int) {
	if currentRow == n {
		*count++
		return
	}

	for i := 0; i < n; i++ {
		idx1 := n - i + currentRow
		idx2 := currentRow + i
		if d1[idx1] || d2[idx2] || columns[i] {
			continue
		}
		d1[idx1], d2[idx2], columns[i] = true, true, true
		backtracking(currentRow+1, n, columns, d1,d2, count)
		d1[idx1], d2[idx2], columns[i] = false, false, false
	}
}
