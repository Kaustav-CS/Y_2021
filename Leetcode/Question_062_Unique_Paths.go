/******************************************************************************
Link:   https://leetcode.com/problems/unique-paths/
62. Unique Paths

Your input
3
7

Output
28

Expected
28

DATE:    2021, October 24
০৭ কার্তিক,১৪২৮
*******************************************************************************/
func uniquePaths(m int, n int) int {
    var cache = make(map[string]int)
	return calculate(m, n, cache)
}

func calculate(m int, n int, cache map[string]int) int {
	var key string = fmt.Sprint(m) + "," + fmt.Sprint(n)
	val, ok := cache[key]
	if ok {
		return val
	} else {
		if (m == 0 || n == 0) {
			return 0
		} else if (m == 1 && n == 1){
			return 1
		} 
		cache[key] = calculate(m-1, n, cache) + calculate(m, n-1, cache)
	    return cache[key]
	}
}
