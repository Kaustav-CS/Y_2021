/******************************************************************************
Link:   https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
116. Populating Next Right Pointers in Each Node

Your input
[1,2,3,4,5,6,7]
Output
[1,#,2,3,#,4,5,6,7,#]
Expected
[1,#,2,3,#,4,5,6,7,#]

DATE:    2021, November 14
২৮ কার্তিক,১৪২৮

*******************************************************************************/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root != nil {
        q := []*Node{root}
        for len(q) > 0 { 
            count := len(q) 
            var prev, n *Node
            for count > 0 {
                n, q = q[0], q[1:]     
                n.Next, prev = prev, n
                if n.Right != nil { 
                    q = append(q, n.Right)
                }
                if n.Left != nil {
                    q = append(q, n.Left)
                }
                count--
            }
        }
    }
    return root
}
