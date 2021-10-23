/******************************************************************************
Link:   https://leetcode.com/problems/rotate-list/
61. Rotate List

Your input
[1,2,3,4,5]
2

Output
[4,5,1,2,3]

Expected
[4,5,1,2,3]

DATE:    2021, October 23
০৬ কার্তিক,১৪২৮

*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil || head.Next == nil { return head }
    dummy, n := &ListNode{-1, head}, 0
    for head != nil { head = head.Next; n++ }
    k = k % n
    if k != 0 {
        a := dummy.Next
        for i := 0; i < n - k - 1; i++ { a = a.Next }
        b := a.Next
        a.Next = nil
        tmp := &ListNode{-1, b}
        for i := 0; i < k - 1; i++ { b = b.Next }
        b.Next = dummy.Next
        return tmp.Next
    }
    return dummy.Next
}
