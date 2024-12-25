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



def solve1(M, guard) -> int:
    """ move the guard until she's out of the map.. """

    visited = set()
    visited.add(guard.position)

    while True:
        next_pos = next_position(guard)
        if next_pos not in M:
            return len(visited)
        elif M[next_pos] == ".":
            guard.position = next_pos
            visited.add(guard.position)
        else:
            # we hit a wall, so we turn the guard 90 degrees, without moving the guard.. 
            guard.direction = turn[guard.direction]
    
    exit("should not reach this")


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
    print(solve1(M, guard))

