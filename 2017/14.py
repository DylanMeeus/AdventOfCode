# Solution to day 14 of Advent Of Code
from collections import OrderedDict
from functools import reduce
import binascii


def solve2():
    # collect the coordinates
    # group the coordinates by scanning one by one, and removing from all coordinates when they are
    # taken

    print("hello world")
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
    coordinates = []
    for row in range(len(bin_values)):
        for column in range(len(bin_values[row])):
            if bin_values[row][column] == '1':
                coordinates.append((row,column))
    groups = []
    while len(coordinates) != 0:
        # take the first, and find all connections
        print(len(coordinates))
        node = coordinates[0]
        group = find_connections(node, [node],coordinates)
        groups.append(group)
        coordinates = list(filter(lambda k: k not in group,coordinates)) 
    print(len(groups))


def find_connections(node, group,coordinates):
    up = (node[0]-1,node[1])
    down = (node[0]+1,node[1])
    left = (node[0],node[1]-1)
    right = (node[0],node[1]+1)
    if up in coordinates and not up in group:
        group.append(up)
        group = find_connections(up,group,coordinates)
    if down in coordinates and not down in group:
        group.append(down)
        group = find_connections(down,group,coordinates)
    if left in coordinates and not left in group:
        group.append(left)
        group = find_connections(left,group,coordinates)
    if right in coordinates and not right in group:
        group.append(right)
        group = find_connections(right,group,coordinates)
    return group


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


#solve()
solve2()
