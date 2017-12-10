# Solution for day 10
from collections import OrderedDict







def getInput():
    inp = [70,66,255,2,48,0,54,48,80,141,244,254,160,108,1,41]
    return inp


def solve():
    inputs = getInput()
    # Create map of indices, makes it easier to move 'm around
    data = {}
    for i in range(256):
        data[i] = i
    current_index = 0
    skip_size = 0
    t = 0
    for value in inputs:
        # take the next X(value) elements
        r_map = OrderedDict()
        print(value)
        for k in range(value):
            r = (current_index + k)%len(data.keys())
            r_map[r] = data[r]

        # Reverse the list
        v = list(r_map.values())
        reversed_list = (v[::-1])
        i = 0
        for r in r_map.keys():
            data[r] = reversed_list[i]
            i += 1
        current_index = (current_index + value + skip_size)%len(data.keys())
        skip_size += 1
    print(data[0] * data[1])
       

        
        

solve()
