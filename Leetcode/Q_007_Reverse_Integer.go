/******************************************************************************

Link:   https://leetcode.com/problems/reverse-integer/
Question_007_Reverse_Integer.go

7. Reverse Integer


Your input
123

Output
321

Expected
321

*******************************************************************************/

import "math"
func reverse(x int) int {
    isPositive := x >= 1
	res := 0
	x = int(math.Abs(float64(x)))

	for x > 0 {
		res = (res * 10) + (x % 10)
		x /= 10
	}
	if res >= math.MaxInt32 {
		return 0
	}
	if !isPositive {
		return -res
	}
	return res
}
