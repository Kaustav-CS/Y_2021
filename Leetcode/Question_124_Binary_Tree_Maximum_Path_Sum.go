/******************************************************************************
Link:   https://leetcode.com/problems/binary-tree-maximum-path-sum/
124. Binary Tree Maximum Path Sum

Your input
[1,2,3]
Output
6
Expected
6

DATE:    2021, November 15
২৯ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    if root == nil{ return 0 }
    res := root.Val
    rec(root, &res)
    return res
}
func rec(node *TreeNode, res *int)int{
    if node == nil{ return 0 }
    l := rec(node.Left, res)
    r := rec(node.Right, res)
    *res = max(*res, node.Val + max(l,0)+ max(r,0))
    return node.Val + max(0, max(r,l))
}
func max (a,b int)int { if a> b { return a }; return b}
