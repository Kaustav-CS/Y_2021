/******************************************************************************
Link:   https://leetcode.com/problems/linked-list-cycle/
141. Linked List Cycle

Your input
[3,2,0,-4]
1

Output
true

Expected
true

DATE:    2021, November 22
০৬ অগ্রহায়ণ,১৪২৮
*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    // Handle edge cases.
	if head == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

