/******************************************************************************
Link:   https://leetcode.com/problems/remove-nth-node-from-end-of-list/
19. Remove Nth Node From End of List


Given an array 'nums' of 'n' integers, return an array of all the 'unique' 
quadruplets '[nums[a], nums[b], nums[c], nums[d]]' such that:

    '0 <= a, b, c, d < n'
    'a, b, c,' and 'd' are 'distinct'.
    'nums[a] + nums[b] + nums[c] + nums[d] == target'

You may return the answer in 'any order'.


Your input
[1,2,3,4,5]
2

Output
[1,2,3,5]

Expected
[1,2,3,5]

DATE:    2021 October 02
১৫ অশ্বিন, ১৪২৮
*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nth := head
    last := head
    for i := 0; i < n+1; i++ {
        if last == nil {
            return head.Next
        }
        last = last.Next
    }

    for last != nil {
        nth = nth.Next        
        last = last.Next        
    }

    nth.Next = nth.Next.Next

    return head
    
}
