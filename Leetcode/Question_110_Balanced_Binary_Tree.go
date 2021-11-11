/******************************************************************************
Link:   https://leetcode.com/problems/balanced-binary-tree/
110. Balanced Binary Tree

Your input
[3,9,20,null,null,15,7]
stdout
3

Output
true

Expected
true

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
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var flag bool = true
    fmt.Println(getHeight(root, &flag))
    return flag
}

func getHeight(root *TreeNode, flag *bool) int {
    if root == nil {
        return 0
    }
    left := getHeight(root.Left, flag)
    right := getHeight(root.Right, flag)
    
    if int(math.Abs(float64(left - right))) > 1{
        fmt.Println(int(math.Abs(float64(left - right))))
        *flag = false
    }
    return int(math.Max(float64(left), float64(right))) + 1
}
