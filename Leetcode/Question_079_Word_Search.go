/******************************************************************************
Link:   https://leetcode.com/problems/word-search/
79. Word Search

Your input
[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
"ABCCED"

Output
true

Expected
true

DATE:    2021, October 29
১২ কার্তিক,১৪২৮

*******************************************************************************/
func exist(board [][]byte, word string) bool {
    for y := 0; y < len(board); y++ {
        for x := 0; x < len(board[y]); x++ {
            if backtrack(board, word, x, y) { return true }
        }
    }

    return false
}

func backtrack(board [][]byte, word string, x, y int) bool {
    var b byte

    if len(word) < 1 { return true }
    if y < 0 || y > len(board)-1 { return false }
    if x < 0 || x > len(board[0])-1 { return false }
    if board[y][x] != word[0] { return false }

    board[y][x] = b

    if backtrack(board, word[1:], x, y-1) { return true }
    if backtrack(board, word[1:], x+1, y) { return true }
    if backtrack(board, word[1:], x, y+1) { return true }
    if backtrack(board, word[1:], x-1, y) { return true }

    board[y][x] = word[0]

    return false
}
