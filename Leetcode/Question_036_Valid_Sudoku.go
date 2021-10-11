/******************************************************************************
Link:   https://leetcode.com/problems/valid-sudoku/
36. Valid Sudoku

Your input
[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]


Output
true


Expected
true

DATE:    2021 October 11
২৪ অশ্বিন, ১৪২৮
*******************************************************************************/
func isValidSudoku(board [][]byte) bool {
    // check squares
    for i := 0; i < 9; i += 3 {
        for j := 0; j < 9; j += 3 {
            if !isValidSubSquare(board, i, j) {
                return false
            }
        }
    }
    
    var rowCheck, colCheck Checker
    // check rows and columns (since width = height, we can do this simultaneously)
    for i := 0; i < 9; i++ {
        rowCheck.Reset()
        colCheck.Reset()
        for j := 0; j < 9; j++ {
            if rowCheck.Seen(board[i][j]) || colCheck.Seen(board[j][i]) {
                return false
            }
            rowCheck.Add(board[i][j])
            colCheck.Add(board[j][i])
        }
    }
	// if everything checks out, return true
    return true
}

// check each 3-3 subsquare
func isValidSubSquare(board [][]byte, x, y int) bool {
    var sqrCheck Checker
    for i := x; i < x+3; i++ {
        for j := y; j < y+3; j++ {
            if sqrCheck.Seen(board[i][j]) {
                return false
            }
            sqrCheck.Add(board[i][j])
        }
    }
    return true
}

// for any given row, column, or 3-3 subgrid you have an array of booleans that store whether or not a given number has already been seen
type Checker struct {
    seen [10]bool
}

// this just resets the array for the next row or column
func (v *Checker) Reset() {
    for i := range v.seen {
        v.seen[i] = false
    }
}

// retrun true if a number has already been used within the row, column, or 3-3 subgrid (making sure to ignore '.' as they are irrelevant to our search)
func (v Checker) Seen(char byte) bool {
    if char == '.' {
        return false
    }
    return v.seen[char - '0']
}

// Add a given number to our seen array so that we will know if this number (gets used more than once
func (v *Checker) Add(char byte) {
    if char == '.' { return }
    v.seen[char - '0'] = true
}
