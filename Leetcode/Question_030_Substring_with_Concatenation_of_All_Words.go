/******************************************************************************
Link:   https://leetcode.com/problems/substring-with-concatenation-of-all-words/
30. Substring with Concatenation of All Words

Your input
"barfoothefoobarman"
["foo","bar"]

Output
[0,9]

Expected
[0,9]

DATE:    2021 October 06
২০ অশ্বিন, ১৪২৮
*******************************************************************************/
func findSubstring(s string, words []string) (ret []int) {
    if len(words) == 0 {return}
    wl := len(words[0])
    for off := 0; off < wl; off++ {
        beg, end := off, off
        ws := append(make([]string, 0), words...)
        loop:
        for end + wl <= len(s) {
            for i, w := range ws {
                if w == s[end:end+wl] {
                    end += wl
                    ws = append(ws[:i], ws[i+1:]...)
                    if len(ws) == 0 {
                        ret = append(ret, beg)
                    }
                    continue loop
                }
            }
            if beg < end {
                ws = append(ws, s[beg:beg+wl])
            } else {
                end += wl
            }
            beg += wl
        }
    }
    return
}
