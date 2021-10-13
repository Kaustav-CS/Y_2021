/******************************************************************************
Link:   https://leetcode.com/problems/count-and-say/
38. Count and Say

Your input
1

Output
"1"

Expected
"1"

DATE:    2021 October 13
২৬ অশ্বিন, ১৪২৮

*******************************************************************************/
func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }    
    s := countAndSay(n-1)    
    count := 0
    result := ""
    var current rune
    for _, c := range s {
        if current != c {              
            result += generate(count, current)
            current = c            
            count = 1             
        } else {
            count += 1
        }        
    }  
    result += generate(count, current)
    return result
}

func generate(count int, current rune) string {
    result := ""
    if count > 0 {
        prepend := fmt.Sprintf("%d", count)
        result += prepend+string(current)
    }
    return result
}
