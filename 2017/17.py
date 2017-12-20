# Solution to spin lock problem
from collections import deque



def solve2():
    print("hello world")
    steps = 369 

    value = 1
    location = 0
    inserts = 1 
    second_position = 0
    for i in range(50000000):
        location = (location + steps) % inserts
        location += 1 # step one more to the right
        if location == 1:
            second_position = value
        value += 1
        inserts += 1
    print(second_position)
    

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
  
#solve()
solve2()
