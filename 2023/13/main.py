from enum import Enum


def lines_to_pattern(lines):
    output = ""
    patterns = []
    for line in lines:
        if line != "":
            output += line + "\n"
        else:
            patterns.append(output)
            output = ""

    return patterns



class axis(Enum):
    HORIZONTAL = 1,
    VERTICAL = 2,




def str_eq(s1, s2):
    shortest = len(s1) if len(s1) < len(s2) else len(s2)

    if shortest == 0:
        return False

    for i in range(0, shortest):
        if s1[i] != s2[i]:
            return False
    return True

def split_string(s, i):
    return [s[0:i+1], s[i+1:]]

def find_vertical_reflection_idx(line):
    reflect_at = []
    for i in range(0, len(line)):
        tpl = split_string(line, i)
        if str_eq(tpl[0][::-1], tpl[1]):
            reflect_at.append(i)
    return reflect_at




def transpose(pattern):
    return [ list(k) for k in zip(*pattern) ] 


def find_reflection(pattern):
    # first find it in the horizontal ones
    lines = pattern.split('\n')[:-1]

    # first find the vertical split
    vert = {}
    for row, line in enumerate(lines):
        if line == "":
            continue
        idx = find_vertical_reflection_idx(line)
        for i in idx:
            if i in vert:
                vert[i] += 1
            else:
                vert[i] = 1

    vert_split = None
    for k, v in vert.items():
        if v == len(lines):
            vert_split = k
    
    # now for the horizontal split
    lines_T = transpose(lines)
    hor = {}
    for row, line in enumerate(lines_T):
        if line == "":
            continue
        idx = find_vertical_reflection_idx(line)
        for i in idx:
            if i in hor:
                hor[i] += 1
            else:
                hor[i] = 1

    hor_split = None
    for k,v in hor.items():
        if v == len(lines_T):
            hor_split = k

    return (vert_split, hor_split)

    



def solve1(patterns):

    _sum = 0
    for pattern in patterns:
        reflection_points = find_reflection(pattern)
        if reflection_points[0] is not None:
            _sum += reflection_points[0] + 1
        else:
            x = (reflection_points[1] + 1) * 100
            _sum += x
    return _sum

        


    return 0

if __name__ == '__main__':
    lines = open('input.txt').read().split("\n")
    patterns = lines_to_pattern(lines)
    print(solve1(patterns))
