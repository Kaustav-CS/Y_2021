/******************************************************************************
Link:   https://leetcode.com/problems/surrounded-regions/
130. Surrounded Regions

Your input
[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]

Output
[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]

Expected
[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]

DATE:    2021, November 17
০১ অগ্রহায়ণ,১৪২৮

*******************************************************************************/
func solve(board [][]byte)  {
   if len(board) == 0 || len(board[0]) == 0 {
        return
    }
    l := len(board)
    ll := len(board[0])
    for i, v := range board[0] {
        if v == 'O' {
            bfs(board, i, 0)
        }
        if board[l - 1][i] == 'O' {
            bfs(board, i, l - 1)
        }
    }
    for i := 0; i < l; i++ {
        if board[i][0] == 'O' {
            bfs(board, 0, i)
        }
        if board[i][ll - 1] == 'O' {
            bfs(board, ll - 1, i)
        }
    }
    for i, v := range board {
        for j, _ := range v {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == 'Y' {
                board[i][j] = 'O'
            }
        }
    }
}

func bfs(m [][]byte, x, y int) {
    m[y][x] = 'Y'
    if x < len(m[0]) - 1 && m[y][x + 1] == 'O' {
        bfs(m, x + 1, y)
    }
    if x > 0 && m[y][x - 1] == 'O' {
        bfs(m, x - 1, y)
    }
    if y < len(m) - 1 && m[y + 1][x] == 'O' {
        bfs(m, x, y + 1)
    }
    if y > 0 && m[y - 1][x] == 'O' {
        bfs(m, x, y - 1)
    }
}

