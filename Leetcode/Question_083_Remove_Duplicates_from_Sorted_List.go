/******************************************************************************
Link:   https://leetcode.com/problems/remove-duplicates-from-sorted-list/
83. Remove Duplicates from Sorted List

Your input
[1,1,2]

Output
[1,2]

Expected
[1,2]


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
    p := head
	for {
		if p == nil || p.Next == nil {
			break
		}
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}
