/******************************************************************************
Link:   https://leetcode.com/problems/4sum/
18. 4Sum


Given an array 'nums' of 'n' integers, return an array of all the 'unique' 
quadruplets '[nums[a], nums[b], nums[c], nums[d]]' such that:

    '0 <= a, b, c, d < n'
    'a, b, c,' and 'd' are 'distinct'.
    'nums[a] + nums[b] + nums[c] + nums[d] == target'

You may return the answer in 'any order'.


Your input
[1,0,-1,0,-2,2]
0

Output
[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Expected
[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

DATE:    2021 October 02
১৫ অশ্বিন, ১৪২৮
*******************************************************************************/
func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
	counts := len(nums)
	var Cnums []int
	var Ctarget int
	var Ccounts int
	var result [][]int
	Rmap := make(map[string]int)

	if counts < 4 {
		return result
	}

	for i:=0;i<counts-3;i++{
		Ctarget = target - nums[i]
		Cnums = nums[i+1:]
		Ccounts = counts - i -1

		for j:=0;j<Ccounts-2;j++{
			k,l := j+1,Ccounts-1
			for ;k<l;{
				if Cnums[j] + Cnums[k] + Cnums[l] == Ctarget{
					Rstring := strconv.Itoa(nums[i]) + strconv.Itoa(Cnums[j]) + strconv.Itoa(Cnums[k]) + strconv.Itoa(Cnums[l])
					if _, ok := Rmap[Rstring]; !ok {
						result = append(result,make4(nums[i],Cnums[j],Cnums[k],Cnums[l]))
						Rmap[Rstring] = 1
					}
					k++
				}

				if Cnums[j] + Cnums[k] + Cnums[l] > Ctarget{
					l--
				}

				if Cnums[j] + Cnums[k] + Cnums[l] < Ctarget{
					k++
				}

			}

		}

	}

	return result
}

func make4(num1 int,num2 int,num3 int, num4 int)[]int{
	Tresult := make([]int,4)
	Tresult[0],Tresult[1],Tresult[2],Tresult[3] = num1,num2,num3,num4
	return Tresult
}
