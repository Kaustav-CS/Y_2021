/******************************************************************************
Link:   https://leetcode.com/problems/minimum-window-substring/
76. Minimum Window Substring

Your input
"ADOBECODEBANC"
"ABC"

Output
"BANC"

Expected
"BANC"

DATE:    2021, October 28
১১ কার্তিক,১৪২৮

*******************************************************************************/
func minWindow(s string, t string) string {
    if s == "" || t == "" {
		return ""
	}

	// t may contains duplicate character, count by map
	member := make(map[rune]int)
	for _, v := range t {
		member[v]++
	}

	// initial minWin, minWindow is s[minWin[0]:minWin[1]+1]
	minWin := [2]int{-1, len(s)-1}
	// l means left bound, consider minWin[0]
	l := 0
	// curMem count character belong s[minWin[0]:minWin[1]+1]
	curMem := make(map[rune]int)
	alreadyOK := false
	for i, v := range s {
		if times, exist := member[v]; exist {
			curMem[v]++
			if curMem[v] >= times {
				if alreadyOK || isOK(member, curMem) {
					alreadyOK = true
					// calculate l here, infinite loop is OK
					for {
						if lt, exist := curMem[rune(s[l])]; exist {
							// if equal, l can't increase any more
							// modify minWin if possible
							if lt == member[rune(s[l])] {
								curWide := i-l+1
								if curWide < minWin[1]-minWin[0]+1 {
									minWin[0] = l
									minWin[1] = i
								}
								break
							}
							// not equal means greater, obviously
							curMem[rune(s[l])]--
						}
						l++
					}
				}
			}
		}
	}
	if minWin[0] == -1 {
		return ""
	}
	return s[minWin[0]:minWin[1]+1]
}

func isOK(standard, now map[rune]int) bool {
	for k, v := range standard {
		if now[k] < v {
			return false
		}
	}
	return true
}
