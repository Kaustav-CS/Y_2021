/******************************************************************************
Link:   https://leetcode.com/problems/merge-intervals/
56. Merge Intervals

Your input
[[1,3],[2,6],[8,10],[15,18]]

Output
[[1,6],[8,10],[15,18]]

Expected
[[1,6],[8,10],[15,18]]


DATE:    2021, October 20
০৩ কার্তিক,১৪২৮
*******************************************************************************/
import "sort"
import "math"

func merge(intervals [][]int) [][]int {
    // ugh. stupid edge case
    n := len(intervals)
    if n == 0 {
        return intervals
    }
    
    var res  [][]int
    sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
    res = append(res, intervals[0])
    
    for i := 1; i < n; i++ {
        lastMerged := res[len(res)-1]
        
        if intervals[i][0] > lastMerged[1] {
            res = append(res, intervals[i])
        } else {
            lastMerged[1] = int(math.Max(float64(intervals[i][1]), float64(lastMerged[1])))
        }
    }
    return res
}
