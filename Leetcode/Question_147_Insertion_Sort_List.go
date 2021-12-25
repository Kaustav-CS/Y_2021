/******************************************************************************
Link:   https://leetcode.com/problems/insertion-sort-list/
147. Insertion Sort List

Your input
[4,2,1,3]

Output
[1,2,3,4]

Expected
[1,2,3,4]

DATE:    2021, December 25
৯ পৌষ,১৪২৮
*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}
	// list[head,..,sortedEnd] is sorted
	sortedEnd := head
	cur := sortedEnd.Next
	var insertAfter *ListNode
	var insertBefore *ListNode
	for cur != nil {
		// don't need to insert
		if cur.Val >= sortedEnd.Val {
			sortedEnd = cur
			cur = cur.Next
		} else {// insert into sorted list
			// expand the sorted list
			sortedEnd.Next = cur.Next
			// initial search
			insertBefore = head
			insertAfter = nil
			for insertBefore != sortedEnd.Next && insertBefore.Val <= cur.Val {
				insertAfter = insertBefore
				insertBefore = insertBefore.Next
			}
			// insert now!
			cur.Next = insertBefore
			// if the cur as the head of list, head must be change
			if insertAfter == nil {
				head = cur
			} else {
				insertAfter.Next = cur
			}
			// next element
			cur = sortedEnd.Next
		}
	}
	return head
}
