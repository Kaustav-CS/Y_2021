/******************************************************************************
Link:   https://leetcode.com/problems/binary-tree-postorder-traversal/
145. Binary Tree Postorder Traversal

Your input
[1,null,2,3]

Output
[3,2,1]

Expected
[3,2,1]

DATE:    2021, November 22
০৬ অগ্রহায়ণ,১৪২৮
*******************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    helper := list.New()
    stack := make([]*TreeNode, 0)
    p := root
    stack = append(stack, p)

    for len(stack) != 0 {
        visit := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        helper.PushFront(visit.Val)
        if visit.Left != nil {
            stack = append(stack, visit.Left)
        }
        if visit.Right != nil {
            stack = append(stack, visit.Right)
        }
    }
    var ret []int
    for helper.Len() > 0 {
        temp := helper.Front()
        ret = append(ret, temp.Value.(int))
        helper.Remove(temp)
    }
    return ret
}

