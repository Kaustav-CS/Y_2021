/******************************************************************************
Link:   https://leetcode.com/problems/permutation-sequence/
60. Permutation Sequence

Your input
3
3

Output
"213"

Expected
"213"

DATE:    2021, October 23
০৬ কার্তিক,১৪২৮

*******************************************************************************/
func getPermutation(n int, k int) string {
    list := []string{}
	for i := 1; i <= n; i++ {
		list = append(list, strconv.Itoa(i))
	}
	k--
	fact := []int{1}
	for i := 1; i <= n; i++ {
		fact = append(fact, fact[len(fact)-1]*i)
	}
	var ans []string
	for i := n; i >= 1; i-- {
		f := fact[i-1]
		y := k / f
		k = k % f
		ans = append(ans, list[y])
		list = remove(list, y)
	}
	return strings.Join(ans, "")
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
