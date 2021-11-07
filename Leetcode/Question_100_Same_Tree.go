/******************************************************************************
Link:   https://leetcode.com/problems/same-tree/
100. Same Tree

Your input
[1,2,3]
[1,2,3]

Output
true

Expected
true

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
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
