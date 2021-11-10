/******************************************************************************
Link:   https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
107. Binary Tree Level Order Traversal II

Your input
[3,9,20,null,null,15,7]

Output
[[15,7],[9,20],[3]]

Expected
[[15,7],[9,20],[3]]

DATE:    2021, November 10
২৪ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
		return nil
	}
	traversal := make([][]int, 0)
	helper(root, 0, &traversal)
	return traversal
}

func helper(root *TreeNode, level int, traversal *[][]int) {
	if root == nil {
		return
	}
	if len(*traversal) < level+1 {
		*traversal = append([][]int{{}}, *traversal...)
	}
	helper(root.Left, level+1, traversal)
	helper(root.Right, level+1, traversal)
	(*traversal)[len(*traversal)-level-1] = append((*traversal)[len(*traversal)-level-1], root.Val)
}
