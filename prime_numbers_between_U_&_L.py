import time
fh = open("Prime_numbers_02.txt",'w')

#L =input('Enter Starting number of the nember set:\t')
L = 10
lower=int(L)

#U =input('Enter Last number of the nember set:\t')
U = 10**31
upper=int(U)


print("Prime numbers between", lower, "and", upper, "are:\t")
print("Prime numbers between", lower, "and", upper, "are:\t",file = fh)

for num in reversed(range(lower, upper + 1)):
   # all prime numbers are greater than 1
   if num > 1:
       for i in range(2, num):
           if (num % i) == 0:
               break
       else:
        print(num )
        print(num, file = fh)
time.sleep(10)
print("\t\t Task End")
quit()
