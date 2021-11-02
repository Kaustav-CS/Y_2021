/******************************************************************************
Link:   https://leetcode.com/problems/unique-binary-search-trees-ii/
95. Unique Binary Search Trees II

Your input
3

Output
[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

Expected
[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

DATE:    2021, November 03
১৭ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 { return []*TreeNode{}}
    return generateTreesRec(1, n)
}
func generateTreesRec(lo, hi int) []*TreeNode {
    res := []*TreeNode{}
    if lo > hi {
        return append(res, nil)
    }
    if lo == hi {
        return append(res, &TreeNode{Val: lo})
    }
    
    
    c := lo
    for c <= hi {
        leftTrees := generateTreesRec(lo, c-1)
        rightTrees := generateTreesRec(c+1, hi)
        
        for _, leftTree := range leftTrees {
            for _, rightTree := range rightTrees {
                res = append(res, &TreeNode{
                    Val: c,
                    Left: leftTree,
                    Right: rightTree,
                })
            }
        }
        c++
    }
    return res
}

