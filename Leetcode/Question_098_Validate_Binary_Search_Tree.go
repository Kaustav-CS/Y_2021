/******************************************************************************
Link:   https://leetcode.com/problems/validate-binary-search-tree/
98. Validate Binary Search Tree

Your input
[2,1,3]

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
func isValidBST(root *TreeNode) bool {
    // time: O(n) where n means the number of nodes
    // space: O(d) where d means the depth of the BST
    return validate(root, nil, nil)
}

func validate(root *TreeNode, max, min *int) bool {
    if root == nil { return true }
    if (max != nil && root.Val >= *max) || (min != nil && root.Val <= *min) { return false }
    
    return validate(root.Left, &(root.Val), min) && validate(root.Right, max, &(root.Val))
}

