/******************************************************************************
Link:   https://leetcode.com/problems/maximum-depth-of-binary-tree/
104. Maximum Depth of Binary Tree

Your input
[3,9,20,null,null,15,7]

Output
3

Expected
3

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
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	
	if l >= r {
		return l + 1
	}
	
	return r + 1
}
