# Python program to show date (YYYY-MM-DD)
#
#
#
#
#
#
#

from datetime import datetime
from time import sleep
f_h = open("Recent_Date_n_Time.txt",'w')
while True:
    now = datetime.now()
    print("Date {},\ttime {}".format(now.strftime("%Y-%m-%d"),now.strftime("%H:%M:%S")))
    print("Date {},\ttime {}".format(now.strftime("%Y-%m-%d"),now.strftime("%H:%M:%S")),file = f_h)
    sleep(4)
