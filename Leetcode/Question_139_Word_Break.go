/******************************************************************************
Link:   https://leetcode.com/problems/copy-list-with-random-pointer/
139. Word Break

Your input
"leetcode"
["leet","code"]

stdout
[-1]
[-1]
[-1]
[-1 3]
[-1 3]
[-1 3]
[-1 3]
[-1 3 7]

Output
true

Expected
true

DATE:    2021, November 19
০৩ অগ্রহায়ণ,১৪২৮

*******************************************************************************/
func wordBreak(s string, wordDict []string) bool {
    checkPoints := make([]int, 0, len(s))
    checkPoints = append(checkPoints, -1)
    
    dict := make(map[string]bool)
    for _, str := range wordDict {
        dict[str] = true
    }
    
    for i:=0; i<len(s); i++ {
        leng := len(checkPoints)
        for j:=1; j<=leng; j++ {
            tail := s[checkPoints[leng-j]+1:i+1]
            if dict[tail] {
                checkPoints = append(checkPoints, i)
                break
            }
        }
        fmt.Println(checkPoints)
    }
    return checkPoints[len(checkPoints)-1] == len(s)-1
}

