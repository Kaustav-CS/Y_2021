/******************************************************************************
Link:   https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
103. Binary Tree Zigzag Level Order Traversal

Your input
[3,9,20,null,null,15,7]

Output
[[3],[20,9],[15,7]]

Expected
[[3],[20,9],[15,7]]

DATE:    2021, November 08
২২ কার্তিক,১৪২৮

*******************************************************************************/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    levels := [][]int{}
    goingLeft := []*TreeNode{}
    goingRight := []*TreeNode{root}
    for {
        // go right
        level := []int{}
        for i := len(goingRight)-1; i >= 0; i-- {
            if goingRight[i] == nil {
                continue
            }
            level = append(level, goingRight[i].Val)
            goingLeft = append(goingLeft, goingRight[i].Left, goingRight[i].Right)
        }
        if len(level) == 0 {
            // nothing left to do
            break
        }
        levels = append(levels, level)
        goingRight = nil
        level = nil
        
        // next level, go left
        for i := len(goingLeft)-1; i >= 0; i-- {
            if goingLeft[i] == nil {
                continue
            }
            level = append(level, goingLeft[i].Val)
            goingRight = append(goingRight, goingLeft[i].Right, goingLeft[i].Left)
        }
        if len(level) == 0 {
            // nothing left to do
            break
        }
        levels = append(levels, level)
        goingLeft = nil
    }
    return levels
}
