/******************************************************************************
Link:   https://leetcode.com/problems/combination-sum/
39. Combination Sum

Your input
[2,3,6,7]
7

Output
[[7],[2,2,3]]

Expected
[[2,2,3],[7]]

DATE:    2021 October 14
২৭ অশ্বিন, ১৪২৮


*******************************************************************************/
func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    sub := []int{}   
    sort.Ints(candidates)
    helper(candidates,target,&result,sub)
    return result
}

func helper(candidates []int,target int,result *[][]int,sub []int) {
    
    if len(candidates)==0 {
        return
    }
    if candidates[0] == target  {
        sub = append(sub, candidates[0]) 
        *result = append(*result, sub)
        return
    } else if candidates[0]<target {       
        helper(candidates[1:],target,result,sub)
        sub2 := make([]int,len(sub))
        copy(sub2,sub)
        sub2 = append(sub2,candidates[0])
        helper(candidates,target-candidates[0],result,sub2)
    } else {
        return
    }
}
