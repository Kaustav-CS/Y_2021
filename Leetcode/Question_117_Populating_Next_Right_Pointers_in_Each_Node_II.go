/******************************************************************************
Link:   https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
117. Populating Next Right Pointers in Each Node II

Your input
[1,2,3,4,5,null,7]

Output
[1,#,2,3,#,4,5,7,#]

Expected
[1,#,2,3,#,4,5,7,#]

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
	if root == nil {
        return nil
    }
    
    if root.Left != nil {
        if root.Right != nil {
            root.Left.Next = root.Right
        } else {
            root.Left.Next = findNext(root)    
        }
    }
    
    if root.Right != nil {
        root.Right.Next = findNext(root)        
    }
    
    connect(root.Right)
    connect(root.Left)
    return root
}


func findNext(root *Node) *Node {
    for ; root != nil; root = root.Next {
        if root.Next != nil {
            if root.Next.Left != nil {
                return root.Next.Left
            } else if root.Next.Right != nil {
                return root.Next.Right
            }
        }
    }

    return nil
}

