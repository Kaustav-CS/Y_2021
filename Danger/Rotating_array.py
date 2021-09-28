f = open("১১_অশ্বিন_১৪২৮.txt",'w')
import time
def permute(nums):
    from itertools import permutations
    n=nums
    l=[]
    perm=permutations(n)
    for i in list(perm):
        l.append(list(i))
    return(l)
nums = []    
for i in range(2**10):
    nums.append(i)
    print(nums)
    time.sleep(0.5)
    #print(permute(nums))
    print(permute(nums),file = f)
    print("\n\n", file = f)
