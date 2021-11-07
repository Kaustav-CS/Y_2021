/******************************************************************************
Link:   https://leetcode.com/problems/symmetric-tree/
101. Symmetric Tree

Your input
[1,2,2,3,4,4,3]

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
func isSymmetric(root *TreeNode) bool {
    bfsQ := []*TreeNode{}
    bfsQ = append(bfsQ,root)
    for len(bfsQ)>0{
        //check palindrome
        for i:=0;i<len(bfsQ)/2;i++ {
            if bfsQ[i] == nil && bfsQ[len(bfsQ)-1-i] == nil {
                continue
            }
            if bfsQ[i] == nil || bfsQ[len(bfsQ)-1-i] == nil {
                return false
            }
            if bfsQ[i].Val != bfsQ[len(bfsQ)-1-i].Val {
                return false
            } 
        }
        for _,n := range bfsQ {
            if n!= nil {
                bfsQ = append(bfsQ,n.Left)
                bfsQ = append(bfsQ,n.Right)            
            }
            bfsQ = bfsQ[1:]
        }
    }
    return true
}
