/******************************************************************************
Link:   https://leetcode.com/problems/merge-two-sorted-lists/
21. Merge Two Sorted Lists

Your input
[1,2,4]
[1,3,4]

Output
[1,1,2,3,4,4]

Expected
[1,1,2,3,4,4]

DATE:    2021 October 03
১৬ অশ্বিন, ১৪২৮
*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var start *ListNode
	var p Pair

	// find the entry node `start`
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val <= l2.Val {
		start = l1
		p.curr = l1
		p.other = l2
	} else {
		start = l2
		p.curr = l2
		p.other = l1
	}

	for p.notDone() {
		p.update()
	}

    // return the entry node `start`
	return start
}

type Pair struct {
	curr  *ListNode
	other *ListNode
}

func (p *Pair) notDone() bool {
	return p.curr != nil && p.other != nil
}

func (p *Pair) update() {
    
	if p.curr.Next == nil {
	    // the current list has run out of nodes
		p.curr.Next = p.other
		p.other = nil
	} else if p.curr.Next.Val <= p.other.Val {
	     // the next item in the `curr` list is smaller than the `other` item
		 // `curr` gets updated to the next item
		p.curr = p.curr.Next
	} else {
	    // the next item in the `curr` list is greater than the other item
		// thus rewire the `curr` node to point to `other`
		// set `other` to `curr.next` and `curr` to `other` 
		p.curr.Next, p.curr, p.other = p.other, p.other, p.curr.Next
	}
}
