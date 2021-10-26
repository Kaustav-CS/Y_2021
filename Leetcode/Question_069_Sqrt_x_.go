/******************************************************************************
Link:   https://leetcode.com/problems/sqrtx/
69. Sqrt(x)

Your input
4

Output
2

Expected
2

DATE:    2021, October 26
০৯ কার্তিক,১৪২৮

*******************************************************************************/
func mySqrt(x int) int {
    c:=0
    i:=1
    for x>0{
        x=x-i
        if x>=0{
             c++
        }
         i=i+2
       
    }
    
    return c
}
