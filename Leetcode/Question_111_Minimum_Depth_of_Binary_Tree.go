/******************************************************************************
Link:   https://leetcode.com/problems/minimum-depth-of-binary-tree/
111. Minimum Depth of Binary Tree

Your input
[3,9,20,null,null,15,7]

Output
2

Expected
2

DATE:    2021, November 11
২৫ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } else if root.Left == nil && root.Right == nil {
        return 1
    } else if root.Left != nil && root.Right == nil {
        leftMinDepthVal := minDepth(root.Left)
        return leftMinDepthVal + 1
    } else if root.Left == nil && root.Right != nil {
        rightMinDepthVal := minDepth(root.Right)
        return rightMinDepthVal + 1
    } else {
        leftMinDepthVal := minDepth(root.Left)
        rightMinDepthVal := minDepth(root.Right)
        if leftMinDepthVal < rightMinDepthVal {
            return leftMinDepthVal + 1
        }
        return rightMinDepthVal + 1
    }
}
