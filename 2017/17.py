# Solution to spin lock problem
from collections import deque


def solve():
    print("hello world")
    steps = 369 

    que = deque()
    que.append(0)
    value = 1
    location = 0
    for i in range(2018):
        location = (location + steps) % len(que)
        location += 1 # step one more to the right
        que.insert(location,value)
        value += 1
    
    for i in range(len(que)):
        if que[i] == 2017:
            print(que[i+1])
            exit()
  
solve()
