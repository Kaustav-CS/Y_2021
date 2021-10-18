/******************************************************************************
Link:   https://leetcode.com/problems/group-anagrams/
49. Group Anagrams

Your input
["eat","tea","tan","ate","nat","bat"]

Output
[["bat"],["eat","tea","ate"],["tan","nat"]]

Expected
[["bat"],["nat","tan"],["ate","eat","tea"]]

DATE:    2021 October 18
০১ কার্তিক,১৪২৮
*******************************************************************************/
func groupAnagrams(strs []string) [][]string {
    ret := make([][]string, 0)
	tmp := make(map[string][]string, 0)
	for _, str := range strs {
		org := str
		ss := strings.Split(str, "")
		sort.Strings(ss)
		str = strings.Join(ss, "")
		tmp[str] = append(tmp[str], org)
	}
	for _, s := range tmp {
		ret = append(ret, s)
	}
	return ret
}
