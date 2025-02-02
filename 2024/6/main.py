import math
from dataclasses import dataclass
from enum import Enum



@dataclass(frozen=True)
class Point:
    x: int
    y: int


class Direction(Enum):
    UP = 1
    RIGHT = 2
    DOWN = 3 
    LEFT = 4


###### maps which I'll use as functions 

char_to_direction = {
        '^': Direction.UP,
        '>': Direction.RIGHT,
        'v': Direction.DOWN,
        '<': Direction.LEFT,
        }

turn = {
        Direction.UP: Direction.RIGHT,
        Direction.RIGHT: Direction.DOWN,
        Direction.DOWN: Direction.LEFT,
        Direction.LEFT: Direction.UP,
        }

###### End of maps which I'll use as functions :) 


@dataclass
class Guard:
    position: Point
    direction: Direction

@dataclass(frozen=True)
class HashableGuard:
    position: Point
    direction: Direction



def get_input():
    """ parse the input as a point-map """ 
    guard = None
    out = {}
    lines = open('input.txt').read().split('\n')
    for row, line in enumerate(lines):
        for col, char in enumerate(line):
            p = Point(row,col)
            out[p] = char
            if char in ['^', '>', '<', 'v']:
                guard_pos = Guard(p, char_to_direction[char])
                out[p] = "."
    return (out, guard_pos)


def next_position(guard) -> Point:
    fm = {
            Direction.UP: Point(guard.position.x - 1, guard.position.y),
            Direction.DOWN : Point(guard.position.x + 1, guard.position.y),
            Direction.LEFT: Point(guard.position.x, guard.position.y - 1),
            Direction.RIGHT: Point(guard.position.x, guard.position.y + 1),
            }
    return fm[guard.direction]



def solve1(M, guard) -> (int, set):
    """ move the guard until she's out of the map.. """

    visited = set()
    visited.add(guard.position)

    while True:
        next_pos = next_position(guard)
        if next_pos not in M:
            return len(visited), visited
        elif M[next_pos] == ".":
            guard.position = next_pos
            visited.add(guard.position)
        else:
            # we hit a wall, so we turn the guard 90 degrees, without moving the guard.. 
            guard.direction = turn[guard.direction]
    
    exit("should not reach this")


def is_looping(M, guard) -> bool:
    visited = set()
    visited.add(HashableGuard(guard.position, guard.direction))

    while True:
        next_pos = next_position(guard)
        if next_pos not in M:
            return False
        if HashableGuard(next_pos, guard.direction) in visited:
            return True
        elif M[next_pos] == ".":
            guard.position = next_pos
            visited.add(HashableGuard(guard.position, guard.direction))
        else:
            # we hit a wall, so we turn the guard 90 degrees, without moving the guard.. 
            guard.direction = turn[guard.direction]
    

def solve2(M, s, guard) -> int:
    result = 0

    m, n = int(math.sqrt(len(M))), int(math.sqrt(len(M)))

    for idx, p in enumerate(s):
        copy_guard = Guard(guard.position, guard.direction)
        M[p] = '#'
        if is_looping(M, copy_guard): 
            result += 1
        M[p] = '.'
        print(f'checked {idx} of {len(s)}')


    return result

def print_map(M, visited = {}):
    for row in range(0,10):
        s = ""
        for col in  range(0, 10):
            p = Point(row,col)
            if p in M and p not in visited:
                s += M[p]
            elif p in visited:
                s += "X"
        print(s)

if __name__ == '__main__':
    M, guard = get_input()

    result, s = (solve1(M, Guard(guard.position, guard.direction)))
    print(result)
    print(solve2(M, s, Guard(guard.position, guard.direction)))

