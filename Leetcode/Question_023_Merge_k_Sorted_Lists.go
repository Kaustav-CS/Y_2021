/******************************************************************************
Link:   https://leetcode.com/problems/merge-k-sorted-lists/
23. Merge k Sorted Lists

Your input
[[1,4,5],[1,3,4],[2,6]]

Output
[1,1,2,3,4,4,5,6]

Expected
[1,1,2,3,4,4,5,6]

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
import "sort"

func mergeKLists(lists []*ListNode) *ListNode {
    m := make(map[int]int)
    e := []int{}
    
    for _, l := range(lists) {
        for l != nil {
            if _, ok := m[l.Val]; ok {
                m[l.Val] += 1
            } else {
                e = append(e, l.Val)
                m[l.Val] = 1
            }
            l = l.Next
        }
    }
    
    dummyHead := &ListNode{Next: nil, Val: 100}
    curNode := dummyHead
    sort.Ints(e)
    for _, v := range(e) {
        for i := 0; i < m[v]; i++ {
            curNode.Next = &ListNode{Next: nil, Val: v}
            curNode = curNode.Next
        }
    }
    
    return dummyHead.Next
}
