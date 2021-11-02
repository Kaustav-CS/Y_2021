/******************************************************************************
Link:   https://leetcode.com/problems/unique-binary-search-trees/
96. Unique Binary Search Trees

Your input
3

Output
5

Expected
5

DATE:    2021, November 03
১৭ কার্তিক,১৪২৮

*******************************************************************************/
func numTrees(n int) int {
    var G []int = make([]int,n+1)
    G[0] = 1
    G[1] = 1
    for i:=2;i<=n;i++ {
        for j:=1;j<=i;j++ {
            G[i] += G[j-1]*G[i-j]
        }
    }
    return G[len(G)-1]
}
