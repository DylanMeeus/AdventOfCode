# Solution to day 14 of Advent Of Code
from collections import OrderedDict
from functools import reduce
import binascii
def solve():
    
    hexbin = {
                '0' : '0000',
                '1' : '0001',
                '2' : '0010',
                '3' : '0011',
                '4' : '0100',
                '5' : '0101',
                '6' : '0110',
                '7' : '0111',
                '8' : '1000',
                '9' : '1001',
                'a' : '1010',
                'b' : '1011',
                'c' : '1100',
                'd' : '1101',
                'e' : '1110',
                'f' : '1111'
            }

    value = "hfdlxzhv"
    values = []
    for i in range(128):
        value_ind = value+"-"+str(i)
        values.append(value_ind)

    values = list(map(lambda k: get_hash(k),values))

    bin_values = []
    for val in values:
        out = reduce(lambda a, b : str(a) + str(b), map(lambda k: hexbin[k],val))
        bin_values.append(out)
    #print(bin_values)
    ones = 0
    for bin_value in bin_values:
        x = len(list(filter(lambda k: k == '1', bin_value)))
        ones += x
    print(ones)
    


def get_hash(data):
    inp = data
    # Conver inputs to their ascii values
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
    for block in range(16):
        sublist = data_values[block_start:block_start+16]
        val = reduce(lambda a,b: a ^ b, sublist)
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
    return out


solve()
