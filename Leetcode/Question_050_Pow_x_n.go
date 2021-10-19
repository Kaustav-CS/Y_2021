/******************************************************************************
Link:   https://leetcode.com/problems/powx-n/
50. Pow(x, n)

Your input
2.00000
10

Output
1024.00000

Expected
1024.00000

DATE:    2021 October 19
০২ কার্তিক,১৪২৮
*******************************************************************************/
func myPow(x float64, n int) float64 {
    switch {
	case n == 0:
		return float64(1)
	case n == 1:
		return x
	case n < 0:
		return 1 / myPow(x, -n)
	}
	if n & 1 == 0 {
		return myPow(x*x, n>>1)
	} else {
		return myPow(x*x, n>>1)*x
	}
}
