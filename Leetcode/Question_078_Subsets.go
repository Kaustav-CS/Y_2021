/******************************************************************************
Link:   https://leetcode.com/problems/subsets/
78. Subsets

Your input
[1,2,3]

Output
[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Expected
[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]


DATE:    2021, October 28
১১ কার্তিক,১৪২৮

*******************************************************************************/
func subsets(nums []int) [][]int {
    var sets [][]int
    for bitmask := uint(0); bitmask < 1 << uint(len(nums)); bitmask++ {
        var set []int
        for index, n := range nums {
            if bitmask & (1 << uint(index)) != 0 {
                set = append(set, n)
            }
        }    
        
        sets = append(sets, set)
    }
    
    return sets
}
