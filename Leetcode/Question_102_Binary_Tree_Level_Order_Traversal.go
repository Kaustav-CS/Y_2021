/******************************************************************************
Link:   https://leetcode.com/problems/binary-tree-level-order-traversal/
102. Binary Tree Level Order Traversal

Your input
[3,9,20,null,null,15,7]

Output
[[3],[9,20],[15,7]]

Expected
[[3],[9,20],[15,7]]

DATE:    2021, November 07
২১ কার্তিক,১৪২৮

*******************************************************************************/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        level := []int{}
        l := len(queue)
        for i := 0; i < l; i++ {
            if queue[0] != nil {
                level = append(level, queue[0].Val)
                queue = append(queue, queue[0].Left)
                queue = append(queue, queue[0].Right)
            }   
            queue = queue[1:]
        }
        result = append(result, level)
    }
    return result[:len(result)-1]
}
