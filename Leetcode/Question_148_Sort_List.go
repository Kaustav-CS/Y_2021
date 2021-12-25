/******************************************************************************
Link:   https://leetcode.com/problems/sort-list/
148. Sort List

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
func sortList(head *ListNode) *ListNode {
    var size int = 0
	for n:=head; n!=nil; n=n.Next {
		size++
	}
	dummy := ListNode{0, head}
	for step:=1; step<size; step*=2 {
		L := &dummy
		for l:=0; l+step<=size; l+=step*2 {
			if l+step*2<=size {
				L = merge(L, step, step)
			} else {
				merge(L, step, size-l-step)
			}
		}
	}
	return dummy.Next
}

func merge(L *ListNode, sizeL, sizeR int) *ListNode {
	R := forwardN(L, sizeL)
	for sizeL>0 && sizeR>0 {
		if L.Next.Val <= R.Next.Val {
			L = L.Next
			sizeL--
		} else {
			r := R.Next
			R.Next = r.Next
			r.Next = L.Next
			L.Next = r
			L = r
			sizeR--
		}
	}
	return forwardN(L, sizeL+sizeR)
}

func forwardN(L *ListNode, n int) *ListNode {
	for ;n>0;n-- {
		L = L.Next
	}
	return L
}
