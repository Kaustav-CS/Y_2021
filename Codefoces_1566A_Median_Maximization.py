#Codefoces_1566A
k2 = []
def code():

    n,s = map(int, input().split())
    #print("\t\t",n,s)
    #s = int(input())
    if(n%2 == 1):
        k = (n//2) +1
    else:
        k = n//2
    q = n-k+1
    z = (s//q)
    k2.append(z)
    #print(z)


t = int(input())
while(t!=0):
    t = t-1
    code()

print("\n")
print("\n")
print("\n")
for j in range(len(k2)):
    print(k2[j])

##
##'''
##        Input
##8
##1 5
##2 5
##3 5
##2 1
##7 17
##4 14
##1 1000000000
##1000000000 1
##        Output
##5
##2
##2
##0
##4
##4
##1000000000
##0
##'''
