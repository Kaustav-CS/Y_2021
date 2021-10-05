/******************************************************************************
Link:   https://leetcode.com/problems/swap-nodes-in-pairs/
24. Swap Nodes in Pairs

Your input
[1,2,3,4]

Output
[2,1,4,3]

Expected
[2,1,4,3]

DATE:    2021 October 05
১৮ অশ্বিন, ১৪২৮
*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode{0, head}
	current := result
	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := first.Next

		temp := second.Next
		second.Next = first
		first.Next = temp
		current.Next = second
		current = current.Next.Next
	}

	return result.Next
}
