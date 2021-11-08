/******************************************************************************
Link:   https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
105. Construct Binary Tree from Preorder and Inorder Traversal

Your input
[3,9,20,15,7]
[9,3,15,20,7]

Output
[3,9,20,null,null,15,7]

Expected
[3,9,20,null,null,15,7]

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
func buildTree(preorder []int, inorder []int) *TreeNode {
    pStart = 0
    return helper(preorder, inorder, 0, len(inorder)-1)
}

func helper(preorder []int, inorder []int, iStart int, iEnd int) *TreeNode {
    if iStart > iEnd {
        return nil
    }
    
    root := &TreeNode{Val: preorder[pStart]}
    pStart++
    
    var idx int
    for idx = iStart; idx <= iEnd; idx++ {
        if root.Val == inorder[idx] {
            break
        }
    }
    
    root.Left = helper(preorder, inorder, iStart, idx-1)
    root.Right = helper(preorder, inorder, idx+1, iEnd)
    
    return root
}
var pStart int
