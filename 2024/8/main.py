import math
from dataclasses import dataclass
from enum import Enum, unique

@dataclass(frozen=True)
class Point:
    row: int
    col: int


@unique
class Direction(Enum):
    UP = 1
    RIGHT_UP = 2
    RIGHT = 3
    RIGHT_DOWN = 4
    DOWN = 5
    LEFT_DOWN = 6
    LEFT = 7
    LEFT_UP = 8



def get_input() -> {}:
    lines = open('input.txt').read().split('\n')
    out = {}
    for row, line in enumerate(lines):
        for col, char in enumerate(line):
            p = Point(row, col)
            out[p] = char
    return out


def print_map(M, antinode_locations = {}):
    N = int(math.sqrt(len(M)))
    for row in range(0, N):
        s = ""
        for col in  range(0, N):
            p = Point(row,col)
            if p in antinode_locations and M[p] == '.':
                s += "#"
            else:
                s += M[p]
        print(s)

def distance(p1, p2) -> int:
    return int(math.sqrt( ((p2.row - p1.row) ** 2) + ((p2.col - p1.col) ** 2 )))

def row_distance(p1, p2) -> int:
    return abs(p1.row - p2.row)

def col_distance(p1, p2) -> int:
    return abs(p1.col - p2.col)

def direction(p1, p2) -> Direction:
    """ returns direction of p1 in relation to p2 """ 
    if p1.row == p2.row and p1.col < p2.row: 
        return Direction.LEFT
    if p1.row == p2.row and p1.col > p2.row: 
        return Direction.RIGHT
    if p1.col == p2.col and p1.col < p2.col:
        return Direction.UP
    if p1.col == p2.col and p1.col > p2.col:
        return Direction.DOWN

    # now the 'sideways positions' 
    if p1.row < p2.row and p1.col < p2.col:
        return Direction.LEFT_UP
    if p1.row < p2.row and p1.col > p2.col:
        return Direction.RIGHT_UP
    if p1.row > p2.row and p1.col < p2.col:
        return Direction.LEFT_DOWN
    if p1.row > p2.row and p1.col > p2.col:
        return Direction.RIGHT_DOWN

    exit('should not reach this')



def antinode_location(start, dsr, dsc, dr, it = 2) -> Point:
    r, c = start.row, start.col
    new_r = r
    new_c = c
    dsr = dsr * it
    dsc = dsc * it
    if dr in [Direction.UP, Direction.RIGHT_UP, Direction.LEFT_UP]:
        new_r = r + dsr 
    if dr in [Direction.DOWN, Direction.RIGHT_DOWN, Direction.LEFT_DOWN]:
        new_r = r - dsr 
    if dr in [Direction.LEFT, Direction.LEFT_DOWN, Direction.LEFT_UP]:
        new_c = c + dsc 
    if dr in [Direction.RIGHT, Direction.RIGHT_DOWN, Direction.RIGHT_UP]:
        new_c = c - dsc 

    return Point(new_r, new_c)



def solve1(data) -> int:
    """ we have to find all positions with antinodes.. """ 
    """ we first have to collect all antennas by their frequency.. """
    freq_map = {}
    for p,f in data.items():
        if f == '.':
            continue
        if f not in freq_map:
            freq_map[f] = []
        freq_map[f].append(p)

    antinodes: [Point] = []

    for freq, points in freq_map.items():
        for point in points:
            for other_point in points:
                if point == other_point:
                    continue
                else:
                    dsr = row_distance(point, other_point)
                    dsc = col_distance(point, other_point) 
                    dr = direction(point, other_point)
                    antinodes.append(antinode_location(point, dsr, dsc, dr))
            

    """ have to filter for ones that are in bounds still """ 
    N = int(math.sqrt(len(data)))
    L = set(filter(lambda k: k.row >= 0 and k.row < N and k.col >= 0 and k.col < N, antinodes))

    return len(L)


def solve2(data) -> int:
    """ we have to find all positions with antinodes.. """ 
    """ we first have to collect all antennas by their frequency.. """
    freq_map = {}
    for p,f in data.items():
        if f == '.':
            continue
        if f not in freq_map:
            freq_map[f] = []
        freq_map[f].append(p)

    antinodes: [Point] = []

    for freq, points in freq_map.items():
        for point in points:
            for other_point in points:
                if point == other_point:
                    continue
                else:
                    antinodes.append(other_point)
                    dsr = row_distance(point, other_point)
                    dsc = col_distance(point, other_point) 
                    dr = direction(point, other_point)
                    anti = antinode_location(point, dsr, dsc, dr)
                    it = 3
                    # as long as the generated point is in the 'map'... 
                    while anti in data: 
                        antinodes.append(anti)
                        anti = antinode_location(point, dsr, dsc, dr, it)
                        it += 1
            

    """ have to filter for ones that are in bounds still """ 
    N = int(math.sqrt(len(data)))

    print_map(data, antinodes)
    
    L = set(filter(lambda k: k.row >= 0 and k.row < N and k.col >= 0 and k.col < N, antinodes))

    return len(L)


if __name__ == '__main__':
    data = get_input()
    print(solve1(data))
    print(solve2(data))
