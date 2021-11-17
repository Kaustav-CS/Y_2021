/******************************************************************************
Link:   https://leetcode.com/problems/clone-graph/
133. Clone Graph

Your input
[[2,4],[1,3],[2,4],[1,3]]

Output
[[2,4],[1,3],[2,4],[1,3]]

Expected
[[2,4],[1,3],[2,4],[1,3]]

DATE:    2021, November 17
০১ অগ্রহায়ণ,১৪২৮
*******************************************************************************/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    visited := make(m)
    return iter(visited, node)
}

func iter(visited m, node *Node) *Node {
    if node == nil {
        return node
    }
    if n, ok := visited[node.Val]; ok {
        return n
    }
    n := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
    visited[n.Val] = n
    for _, v := range node.Neighbors {
        nn := iter(visited, v)
        n.Neighbors = append(n.Neighbors, nn)
    }
    return n
}
type m map[int]*Node

