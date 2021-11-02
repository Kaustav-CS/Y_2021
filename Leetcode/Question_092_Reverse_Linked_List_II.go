/******************************************************************************
Link:   https://leetcode.com/problems/reverse-linked-list-ii/
92. Reverse Linked List II

Your input
[1,2,3,4,5]
2
4

Output
[1,4,3,2,5]

Expected
[1,4,3,2,5]

DATE:    2021, November 02
১৬ কার্তিক,১৪২৮
*******************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    arr := []int{}
    listToArr(head, &arr)
    
    start, end := min0(left-1), min0(right-1)
    
    temp := append(arr[:start], reverse(arr[start:end+1])...)
    temp = append(temp, arr[end+1:]...)
    
    resHead := NewList(0)
    arrToList(temp, 0, resHead)
    
    return resHead.Next
}

func min0(x int) int {
    if x < 0 {
        return 0
    }
    return x
}

func arrToList(arr []int, idx int, curr *ListNode) {
    if idx >= len(arr) {
        return
    }
    
    curr.Next = NewList(arr[idx])
    curr = curr.Next

    idx++
    
    arrToList(arr, idx, curr)
}

func NewList(val int) *ListNode {
    return &ListNode{
        Val: val,
    }
}

func listToArr(head *ListNode, arr *[]int) {
    if head == nil {
        return
    }
    *arr = append(*arr, head.Val)
    
    listToArr(head.Next, arr)
}

func reverse(arr []int) []int {
    res := make([]int, 0)
    for i := len(arr)-1; i >= 0; i-- {
        res = append(res, arr[i])
    }
    return res
}
