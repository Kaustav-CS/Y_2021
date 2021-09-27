/******************************************************************************
Link:   https://leetcode.com/problems/regular-expression-matching/

10. Regular Expression Matching

Your input
"aa"
"a"

Output
false

Expected
false

DATE:    2021 September 27
১০ অশ্বিন, ১৪২৮
*******************************************************************************/
func isMatch(s string, p string) bool {
    re := regexp.MustCompile(fmt.Sprintf("^%s$", p))
	return re.Match([]byte(s))
}
