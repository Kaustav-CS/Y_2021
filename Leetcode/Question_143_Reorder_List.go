/******************************************************************************
Link:   https://leetcode.com/problems/reorder-list/
143. Reorder List

Your input
[1,2,3,4]

Output
[1,4,2,3]

Expected
[1,4,2,3]

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
func reorderList(head *ListNode)  {
    if head == nil {
		return
	}
	p1 := head
	p2 := head.Next
	// Find the intermediate node
	for ; p2 != nil && p2.Next != nil; {
		p2, p1 = p2.Next.Next, p1.Next
	}

	head2 := p1.Next
	p1.Next = nil

	head2 = reverseList(head2)

	pre := new(ListNode)
	x1, x2 := head, head2
	for ; x1 != nil && x2 != nil; {
		temp := x1.Next
		x1.Next = x2
		pre.Next = x1
		pre = x2

		x1 = temp
		x2 = x2.Next

	}
	if x1 != nil {
		pre.Next = x1
	}

}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil; {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre

}

