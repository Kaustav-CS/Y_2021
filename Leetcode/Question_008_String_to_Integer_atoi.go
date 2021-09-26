/******************************************************************************

Link:   https://leetcode.com/problems/string-to-integer-atoi/
Question_008_String_to_Integer_atoi.go

8. String to Integer (atoi)


Your input
"42"

Output
42

Expected
42
*******************************************************************************/
func myAtoi(s string) int {
    const MaxInt32 = int(2147483647)
const MinInt32 = int(-2147483648)
    var result int
	var isFound bool

	isNegative := 1

	for i:=0; i<len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
            isFound = true
			v := int(s[i] - '0')
			if result * 10 > MaxInt32- v {
				result = MaxInt32
			}
			result *= 10
			result += v
			continue
		}
		if s[i] == ' ' && !isFound {
			continue
		}
		if s[i] == '-' ||  s[i] == '+'{
			if !isFound {
				isFound = true
				if s[i] == '-' {
					isNegative = -1
				}
				continue
			}

			break
		}

		if !isFound {
			result = 0
		}

		break
	}

	result *= isNegative

	if MaxInt32 < result {
		return MaxInt32
	}

	if MinInt32 > result {
		return MinInt32
	}

	return int(result)
}
