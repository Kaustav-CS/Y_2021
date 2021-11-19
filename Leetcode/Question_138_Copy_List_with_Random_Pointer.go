/******************************************************************************
Link:   https://leetcode.com/problems/copy-list-with-random-pointer/
138. Copy List with Random Pointer

Your input
[[7,null],[13,0],[11,4],[10,2],[1,0]]

Output
[[7,null],[13,0],[11,4],[10,2],[1,0]]

Expected
[[7,null],[13,0],[11,4],[10,2],[1,0]]

DATE:    2021, November 19
০৩ অগ্রহায়ণ,১৪২৮

*******************************************************************************/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    m := make(map[*Node]*Node)
    
    dummy := &Node{}
    
    curr1 := head
    curr2 := dummy
    for curr1 != nil {
        newNode := &Node{Val: curr1.Val}
        m[curr1] = newNode
        curr2.Next = newNode
        curr1 = curr1.Next
        curr2 = curr2.Next
    }
    
    curr1 = head
    curr2 = dummy.Next
    for curr1 != nil {
        curr2.Random = m[curr1.Random]
        curr1 = curr1.Next
        curr2 = curr2.Next
    }
    
    return dummy.Next
}
