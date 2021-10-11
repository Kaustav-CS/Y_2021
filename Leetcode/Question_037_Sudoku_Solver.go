/******************************************************************************
Link:   https://leetcode.com/problems/sudoku-solver/
37. Sudoku Solver

Your input
[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]

Output
[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]

Expected
[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]


DATE:    2021 October 11
২৪ অশ্বিন, ১৪২৮
*******************************************************************************/
var entries = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
func solveSudoku(board [][]byte)  {
    // an extra boolean pointer to stop the recursion
	var t bool = false
	finish := &t
	solver(board, finish)
}

func solver(board [][]byte, finish *bool) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == '.' {
				for k := 0; k < 9; k++ {
					n := entries[k]
					if possible(x, y, n, board) {
						board[x][y] = n  // set it as tmp answer
						solver(board, finish)  // keep solving the remaining board
						if !*finish {
						    // the tmp anwser is not correct, we return to here and mark it as empty again for next try
							board[x][y] = '.'
						}
					}
				}
				// if all number are not fit the current empty cell,
				// there is some wrong anwser, so we return
				return 
			}
		}
	}
	// if no empty cell and we reach here, we have finished solving the board
	// if we don't set this flag, the solver will return to last call and mark the cell empty again and again
	// and eventually makes the board become inital state again
	*finish = true
	return
}

// for checking the move valid or not
func possible(x, y int, n byte, board [][]byte) bool {
	// check horizontal
	for k := 0; k < 9; k++ {
		if board[x][k] == n {
			return false
		}
	}
	// check veritical
	for k := 0; k < 9; k++ {
		if board[k][y] == n {
			return false
		}
	}
	// check small square
	// sqX, sqY = 0/1/2
	sqX := x / 3
	sqY := y / 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[3*sqX+i][3*sqY+j] == n {
				return false
			}
		}
	}

	return true
}
