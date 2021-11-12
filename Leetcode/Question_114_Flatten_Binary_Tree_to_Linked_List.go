/******************************************************************************
Link:   https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
114. Flatten Binary Tree to Linked List

Your input
[1,2,5,3,4,null,6]

Output
[1,null,2,null,3,null,4,null,5,null,6]

Expected
[1,null,2,null,3,null,4,null,5,null,6]


DATE:    2021, November 12
২৬ কার্তিক,১৪২৮
*******************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    for root != nil {
		if root.Left != nil && root.Right != nil{
			t := root.Left
			for t.Right != nil {
				t = t.Right
			}
			t.Right = root.Right
		}
		
		if root.Left != nil {
			root.Right = root.Left
		}
		root.Left = nil
		root = root.Right
	}
}
