/******************************************************************************
Link:   https://leetcode.com/problems/path-sum/
112. Path Sum

Your input
[5,4,8,11,null,13,4,7,2,null,null,null,1]
22

Output
true

Expected
true

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
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {return false}
    
    if (root.Right == nil && root.Left == nil) {
        if root.Val == targetSum {return true}else{return false}
    }
    
    return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
    
}
