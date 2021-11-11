/******************************************************************************
Link:   https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
109. Convert Sorted List to Binary Search Tree

Your input
[-10,-3,0,5,9]

Output
[0,-3,9,-10,null,5]

Expected
[0,-3,9,-10,null,5]

DATE:    2021, November 11
২৫ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    
    dummy := new(ListNode)
    dummy.Next = head
    slow, fast := dummy, dummy
    
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    mid := slow.Next
    
	// cut the tail of the Left, dummy.Next-> ... -> slow -> mid -> ...
	slow.Next = nil
	
    return &TreeNode{
        mid.Val, 
        sortedListToBST(dummy.Next),
        sortedListToBST(mid.Next),
    }
}
