/******************************************************************************
Link:   https://leetcode.com/problems/longest-common-prefix/
14. Longest Common Prefix

Your input
["flower","flow","flight"]

Output
"fl"

Expected
"fl"

DATE:    2021 September 30
১৩ অশ্বিন, ১৪২৮
*******************************************************************************/
func longestCommonPrefix(strs []string) string {
    var vals string
	if len(strs) == 0 {
		return ""
	}
	vals = strs[0]
	for _, stringValue := range strs {
		vals = compareAndReturn(vals, stringValue)
	}
	return vals
}

func compareAndReturn(firstString, secondString string) string {
	var result []byte
	for i := 0; i < len(firstString) && i < len(secondString); i++ {
		if firstString[i] != secondString[i] {
			break
		}
		result = append(result, firstString[i])
	}
	return string(result)
}
