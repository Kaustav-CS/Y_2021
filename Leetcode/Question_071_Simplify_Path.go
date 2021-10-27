/******************************************************************************
Link:   https://leetcode.com/problems/simplify-path/
71. Simplify Path

Your input
"/home/"

Output
"/home"

Expected
"/home"

DATE:    2021, October 27
১০ কার্তিক,১৪২৮
*******************************************************************************/
func simplifyPath(path string) string {
    var stack []string
	for _, str := range strings.Split(path, "/") {
		switch str {
		case "", ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, str)
		}
	}
	return "/" + strings.Join(stack, "/")
}
