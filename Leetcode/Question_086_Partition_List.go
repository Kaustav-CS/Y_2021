/******************************************************************************
Link:   https://leetcode.com/problems/partition-list/
86. Partition List

Your input
[1,4,3,2,5,2]
3

Output
[1,2,2,4,3,5]

Expected
[1,2,2,4,3,5]

DATE:    2021, October 31
১৪ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    leftHead := &ListNode{}
    rightHead := &ListNode{}
    
    left := leftHead
    right := rightHead
    for head != nil {
        if head.Val < x {
            left.Next = head
            left = left.Next
        } else {
            right.Next = head
            right = right.Next
        }
        head = head.Next
    }
    
    right.Next = nil
    left.Next = rightHead.Next
    return leftHead.Next
}
