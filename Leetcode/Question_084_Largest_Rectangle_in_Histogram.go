/******************************************************************************
Link:   https://leetcode.com/problems/largest-rectangle-in-histogram/
84. Largest Rectangle in Histogram

Your input
[2,1,5,6,2,3]
Output
10
Expected
10


DATE:    2021, October 30
১৩ কার্তিক,১৪২৮
*******************************************************************************/
func largestRectangleArea(heights []int) int {
    i,max := 0, 0
    stack := make([]int,0)
    
    for i < len(heights){
        if (len(stack) == 0 || heights[i] > heights[stack[len(stack) -1]]){
            stack = append(stack, i);
            i++
        } else{
            pop := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            h := heights[pop]
            var w int 
            if len(stack) == 0 {
                w = i
            } else{
                w = i - stack[len(stack)-1] -1
            }
            max = findMax(max, h * w)
        }
    }
    
    for len(stack) != 0{
        pop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        h := heights[pop]
        var w int 
        if len(stack) == 0 {
            w = i
        } else{
            w = i - stack[len(stack)-1] -1
        }
        max = findMax(max, h * w)
    }
    return max
}

func findMax(a, b int) int{
    if a > b{
        return a
    }
    return b
}
