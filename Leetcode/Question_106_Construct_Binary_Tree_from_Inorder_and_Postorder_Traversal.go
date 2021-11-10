/******************************************************************************
Link:   https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
106. Construct Binary Tree from Inorder and Postorder Traversal

Your input
[9,3,15,20,7]
[9,15,7,20,3]

Output
[3,9,20,null,null,15,7]

Expected
[3,9,20,null,null,15,7]

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
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	rootEl := &TreeNode{Val: rootVal}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootEl.Left = buildTree(inorder[:i], postorder[:i])
			rootEl.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
		}
	}
	return rootEl
}
