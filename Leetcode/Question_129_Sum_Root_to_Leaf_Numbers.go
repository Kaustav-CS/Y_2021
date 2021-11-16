/******************************************************************************
Link:   https://leetcode.com/problems/sum-root-to-leaf-numbers/
129. Sum Root to Leaf Numbers

Your input
[1,2,3]

Output
25

Expected
25

DATE:    2021, November 16
৩০ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    sum := 0
    helper(root, 0, &sum)
    return sum
}

func helper(root *TreeNode, currentSum int, sum *int) {
    if root == nil {
       return 
    }
    
    currentSum = currentSum * 10 + root.Val
    if root.Left == nil && root.Right == nil {
        *sum += currentSum
        return
    }
    
    helper(root.Left, currentSum, sum)
    helper(root.Right, currentSum, sum)
}

