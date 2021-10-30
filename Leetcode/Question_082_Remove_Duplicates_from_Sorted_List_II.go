/******************************************************************************
Link:   https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
82. Remove Duplicates from Sorted List II

Your input
[1,2,3,3,4,4,5]

Output
[1,2,5]

Expected
[1,2,5]

DATE:    2021, October 30
১৩ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    fakehead := new(ListNode)
    fakehead.Next = head
    pre := fakehead
    cur := head
    for cur != nil {
        for cur.Next != nil && cur.Val == cur.Next.Val {
            cur = cur.Next
        }
        if pre.Next == cur {
            pre = pre.Next
        } else {
            pre.Next = cur.Next
        }
        cur = cur.Next
    }
    return fakehead.Next
}
