#N = int(input())
N = int(11)
M = 3*N
#N, M = map(int,input().split())
# M = 3N
# N =9, M = 27
# N = 7, M = 21
# N = 11, M = 33
# 5 < N < 101
# 15 < M < 303
fh = open("Design_001_০২_আশ্বিন_১৪২৮.txt",'w')
for i in range(1,N,2):
    print((i * ".|.").center(M, "-"))
    print((i * ".|.").center(M, "-"), file = fh)
print("WELCOME".center(M,"-"))
print("  Reduce Use of Electricity  ".center(M,"-"), file = fh)
print("  GO TO SLEEP  ".center(M,"-"), file = fh)
for i in range(N-2,-1,-2):
    print((i * ".|.").center(M, "-"))
    print((i * ".|.").center(M, "-"), file = fh)
