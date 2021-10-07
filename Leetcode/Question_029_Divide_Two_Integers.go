/******************************************************************************
Link:   https://leetcode.com/problems/divide-two-integers/
29. Divide Two Integers
Your input
10
3
Output
3
Expected
3

DATE:    2021 October 06
২০ অশ্বিন, ১৪২৮
*******************************************************************************/
func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	var sign int
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	} else {
		sign = -1
	}

	result, power := uint(0), uint(32)
	dividend, divisor = abs(dividend), abs(divisor)
	divPower := uint(divisor) << power

	for dividend >= divisor {
		for divPower > uint(dividend) {
			divPower >>= 1
			power--
		}
		result += 1 << power
		dividend -= int(divPower)
	}

	return int(result) * sign
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
