/******************************************************************************
Link:   https://leetcode.com/problems/reverse-nodes-in-k-group/
25. Reverse Nodes in k-Group

Your input
[1,2,3,4,5]
2

Output
[2,1,4,3,5]

Expected
[2,1,4,3,5]

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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}

	shadowHead := &ListNode{
		Next: head,
	}

	a := shadowHead
	b := a.Next
	c := b
	loop:
	for {
		for i := 0; i < k; i++ {
			if c == nil {
				break loop
			}
			if i != k-1 {
				c = c.Next
			}
		}
		d := c.Next // d could be nil
		// a->b->b1...bN->c->d  ===>  a-><-b<-b1...bN<-c d
		prev, cur := a, b
		for i := 0; i < k; i++ {
			nextCur := cur.Next
			nextPrev := cur
			cur.Next = prev
			cur = nextCur
			prev = nextPrev
		}

		// a-><-b<-b1...bN<-c d  ===>  a->c->bN...b1->b->d
		a.Next = c
		b.Next = d

		// assign a,b,c for next loop
		a = b
		b = d
		c = b
	}
	return shadowHead.Next
}


