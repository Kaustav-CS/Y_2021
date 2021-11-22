/******************************************************************************
Link:   https://leetcode.com/problems/binary-tree-preorder-traversal/
144. Binary Tree Preorder Traversal

Your input
[1,null,2,3]

Output
[1,2,3]

Expected
[1,2,3]

DATE:    2021, November 22
০৬ অগ্রহায়ণ,১৪২৮
*******************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    r := make([]int, 0)
	if root == nil {
		return r
	}
	stack := StackNode{}
	stack.Push(root)
	for !stack.IsEmpty() {
		cur := stack.Pop()
		r = append(r, cur.Val)
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}
	return r
}

type StackNode struct {
	data [1000]*TreeNode
	pos  int
}

func NewStackNode() *StackNode {
	return &StackNode{}
}
func (s *StackNode) IsEmpty() bool {
	return s.pos == 0
}
func (s *StackNode) Push(n *TreeNode) {
	s.data[s.pos] = n
	s.pos++
}
func (s *StackNode) Pop() (n *TreeNode) {
	s.pos--
	return s.data[s.pos]
}

