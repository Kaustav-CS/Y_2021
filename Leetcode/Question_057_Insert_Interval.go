/******************************************************************************
Link:   https://leetcode.com/problems/insert-interval/
57. Insert Interval

Your input
[[1,3],[6,9]]
[2,5]

Output
[[1,5],[6,9]]

Expected
[[1,5],[6,9]]


DATE:    2021, October 21
০৪ কার্তিক,১৪২৮

*******************************************************************************/
func insert(intervals [][]int, newInterval []int) [][]int {
    result := [][]int{}

    if len(intervals) == 0 || len(intervals[0]) == 0 {
        
        if len(newInterval) != 0 {
            result = append(result, newInterval)
        }
        
        return result
    }
    
    newStart := newInterval[0]
    newEnd := newInterval[1]
    i := 0
    
    for i < len(intervals) && newStart > intervals[i][1] {
        result = append(result, intervals[i])
        i++
    } 
    
    start := newStart
    end := newEnd
    
    for i < len(intervals) && newEnd >= intervals[i][0] {
        start = min(start, min(newStart, intervals[i][0]))
        end = max(end, max(newEnd, intervals[i][1]))
        i++
    }
    result = append(result, []int{start, end})
    
    for i < len(intervals) {
        result = append(result, intervals[i])
        i++
    }
    
    return result
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
