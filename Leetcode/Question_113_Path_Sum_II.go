/******************************************************************************
Link:   https://leetcode.com/problems/path-sum-ii/
113. Path Sum II

Your input
[5,4,8,11,null,13,4,7,2,null,null,5,1]
22

Output
[[5,4,11,2],[5,8,4,5]]

Expected
[[5,4,11,2],[5,8,4,5]]


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
func pathSum(root *TreeNode, targetSum int) [][]int {
    res := [][]int{}
    ans := []int{}
    var recurse func(*TreeNode, int)
    recurse = func(r *TreeNode, target int){
        if r == nil {
            return
        }
        ans = append(ans, r.Val)
        target = target - r.Val
        if target == 0 {
            if r.Left == nil && r.Right == nil {
                res = append(res, ans)
                ans = ans[:len(ans)-1:len(ans)-1]
                return
            }
        }
        recurse(r.Left, target)
        recurse(r.Right, target)
        ans = ans[:len(ans)-1:len(ans)-1]
        
    }
    recurse(root, targetSum)
    return res
}
