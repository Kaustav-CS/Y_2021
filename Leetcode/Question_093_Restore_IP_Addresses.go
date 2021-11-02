/******************************************************************************
Link:   https://leetcode.com/problems/restore-ip-addresses/
93. Restore IP Addresses

Your input
"25525511135"

Output
["255.255.11.135","255.255.111.35"]

Expected
["255.255.11.135","255.255.111.35"]

DATE:    2021, November 02
১৬ কার্তিক,১৪২৮
*******************************************************************************/
func restoreIpAddresses(s string) []string {
    var buffer []rune
    var res []string
	// 0 is the current index, 1 is the current level
	restore(s, 0, 1, buffer, &res)
	return res
}

func restore(s string, i, l int, buffer []rune, res *[]string) {
	n := len(s)
	// if level is bigger than 4, then the current result  already has four parts
	if l > 4 {
	// if current index equals the length of s, it means all characters in s has been added to the IP address and the IP address is valid,
	// then we can append current result to the final result
		if i == n {
			buffer = buffer[:len(buffer)-1]
			cur := make([]rune, len(buffer))
			copy(cur, buffer)
			*res = append(*res, string(cur))
		}
		return
	}
	var num int
	// can't have leading 0
	if i < n && s[i] == '0' {
		buffer = append(buffer, '0', '.')
		restore(s, i+1, l+1, buffer,res)
	} else {
	//from current i, we start to append the next 1, 2 and 3 numbers to the current result
		for j := i; j < i+3 && j < n; j++ {
			num = 10*num + int(s[j]-'0')
			//the number has to be less than 255
			if num <= 255 {
			//append number at index i, from index i to i+1, from i to i+2 to the current result
				buffer = append(buffer, []rune(s[i:j+1])...)
				//append "." to separate
				buffer = append(buffer, '.')
				restore(s, j+1, l+1, buffer,res)
				//truncate the last few characters appended before. j-i+1 is the number of characters appended before, the -1 is for the "."
				buffer = buffer[:len(buffer)-(j-i+1)-1] 
			}
		}
	}
}
