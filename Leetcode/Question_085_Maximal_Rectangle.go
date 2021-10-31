/******************************************************************************
Link:   https://leetcode.com/problems/maximal-rectangle/
85. Maximal Rectangle

Your input
[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output
6
Expected
6


DATE:    2021, October 31
১৪ কার্তিক,১৪২৮

*******************************************************************************/
func maximalRectangle(matrix [][]byte) int {
    row:= len(matrix)
    if row == 0 {
        return 0
    }
    col:= len(matrix[0])
    
    arr := make([]int, col)
    max := 0
    for i:=0; i<row; i++{
        for j:=0; j<col; j++{
            cur:= int(matrix[i][j] -'0')
            if cur == 0{
                arr[j] = 0
            } else{
                arr[j] = arr[j] + cur
            }
        }
        area := maxRect(arr)
        if  area > max {
            max = area
        }
    }
    return max
}

func maxRect (arr []int) int{
    var stack []int
    max := 0
    for i:=0; i<=len(arr); i++{
        
        var cur int
        if i == len(arr){
            cur = -1
        } else{
            cur = arr[i]
        }
       
        for (len(stack)!=0 && cur < arr[stack[len(stack)-1]]) {
            width := arr[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            var height int
            if len(stack) == 0{
                height =i
            } else{
                height = i- stack[len(stack)-1] -1
            }
            area := width * height
            if area > max {
                max = area
            }
        }
        stack = append(stack, i)
    }
    return max
    
}
