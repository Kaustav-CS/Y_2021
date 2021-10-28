/******************************************************************************
Link:   https://leetcode.com/problems/search-a-2d-matrix/
74. Search a 2D Matrix

Your input
[[1,3,5,7],[10,11,16,20],[23,30,34,60]]
3

Output
true

Expected
true

DATE:    2021, October 28
১১ কার্তিক,১৪২৮

*******************************************************************************/
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    if n == 0 {
        return false
    }
    
    left, right := 0, m - 1
    
    for left <= right {
        mid := (right - left) >> 1 + left
        if target >= matrix[mid][0] && target <= matrix[mid][n - 1] {
            p := 0
            q := n - 1
            for p <= q {
                if target == matrix[mid][(q - p) >> 1 + p] {
                    return true
                } else if target < matrix[mid][(q - p) >> 1 + p] {
                    q = (q - p) >> 1 + p - 1
                } else {
                    p = (q - p) >> 1 + p + 1
                }
            }
            return false
        } else if target < matrix[mid][0] {
            right = mid - 1
        } else if target > matrix[mid][n - 1] {
            left = mid + 1
        }
    }   
    return false
}
