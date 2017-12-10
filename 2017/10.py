# Solution for day 10
from collections import OrderedDict
from functools import reduce






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
       


def solve2():
    inputs = getInput()
    inp = "70,66,255,2,48,0,54,48,80,141,244,254,160,108,1,41"
    # Convert inputs to their ascii values
    format_in = []
    result = map(lambda y: ord(y), inp)
    inputs = (list(result))
    inputs.extend((17,31,73,47,23))


    # Create map of indices, makes it easier to move 'm around
    data = {}
    for i in range(256):
        data[i] = i
    current_index = 0
    skip_size = 0
    t = 0
    for cycle in range(64):
        for value in inputs:
            # take the next X(value) elements
            r_map = OrderedDict()
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
    # create dense hash of 16
    data_values = list(data.values())
    block_results = []
    block_start = 0
    print(data_values)
    for block in range(16):
        sublist = data_values[block_start:block_start+16]
        val = reduce(lambda a,b: a ^ b, sublist)
        print(sublist)
        block_results.append(val)
        block_start += 16

    hex_nums = list(map(lambda l: hex(l), block_results))
    hex_nums = list(map(lambda l: str(l)[2:],hex_nums))
    out = ""
    for num in hex_nums:
        if len(num) == 1:
            num = '0'+num
        out += num
    print(out)
    

   
    
        

solve2()
