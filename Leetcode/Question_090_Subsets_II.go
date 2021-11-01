/******************************************************************************
Link:   https://leetcode.com/problems/subsets-ii/
90. Subsets II

Your input
[1,2,2]

Output
[[],[1],[1,2],[1,2,2],[2],[2,2]]

Expected
[[],[1],[1,2],[1,2,2],[2],[2,2]]

DATE:    2021, November 01
১৫ কার্তিক,১৪২৮

*******************************************************************************/
import "sort"
func subsetsWithDup(nums []int) [][]int {
    ch := make(chan []int)
    
    sort.Ints(nums)
    
    go func() {
        ch <- []int{}
        generate([]int{}, nums, ch)
        
        close(ch)
    }()
    
    result := [][]int{}
    
    for x := range ch {
        result = append(result, x)
    }
    
    return result
}

func generate(s []int, nums []int, ch chan []int) {
    prev := (1 << 31) - 1
    n := len(s) + 1
    m := len(nums)
    
    if len(nums) == 0 {
        return
    }
    
    for i, x := range nums {
        if x == prev {
            continue
        }
        
        prev = x
        
        dst := make([]int, n - 1, n)
        copy(dst, s)
        dst = append(dst, x)
        
        ch <- dst
        
        if i < m - 1 {
            generate(dst, nums[i+1:], ch)
        }
    }
}

