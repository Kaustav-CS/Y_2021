/******************************************************************************
Link:   https://leetcode.com/problems/linked-list-cycle-ii/
142. Linked List Cycle II

Your input
[3,2,0,-4]
1

Output
tail connects to node index 1

Expected
tail connects to node index 1


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
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}
	cycleNode := make(map[*ListNode]int)
	first := head
	fast := head.Next
	slow := head
    isCycle := false
	for fast != nil && fast.Next != nil {
		if fast == slow {
			isCycle = true
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	for isCycle {
		if _, ok := cycleNode[slow]; !ok {
			cycleNode[slow]++
			slow = slow.Next
		} else {
			break
		}
	}
	for first != nil   {
		if _, ok := cycleNode[first]; ok {
			return first
		}
		first = first.Next
	}
	return nil
}

