/******************************************************************************
Link:   https://leetcode.com/problems/add-two-numbers/
Question_002_Add_Two_Numbers


''
You are given two 'non-empty' linked lists representing two non-negative integers.
The digits are stored in 'reverse order', and each of their nodes contains a
single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the
number 0 itself.

Your input
[2,4,3]
[5,6,4]

Output
[7,0,8]

Expected
[7,0,8]

*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * 
     
 }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
 dummy := ListNode{}
    carry, sum := 0, 0
    p, p1, p2 := &dummy, l1, l2
    
    for p1 != nil || p2 != nil || carry != 0 {
        sum = carry
        if p1 != nil {
            sum += p1.Val
            p1 = p1.Next
        }
        if p2 != nil {
            sum += p2.Val
            p2 = p2.Next
        }
        carry = sum / 10
        p.Next = &ListNode{Val: sum % 10, Next: nil}
        p = p.Next
    }
    
    return dummy.Next   
}
