from dataclasses import dataclass
from typing import Dict, Tuple, Set, List
import copy 


@dataclass(frozen=True)
class Point:
    row: int
    col: int


def get_input() -> Tuple[Dict[Point, int], Set[Point]]:
    lines = open('input.txt').read().split('\n')
    out = {}
    zeroes = set()
    for row, line in enumerate(lines):
        for col, height in enumerate(line):
            p = Point(row,col)
            out[p] = int(height)
            if int(height) == 0 :
                zeroes.add(p)
    return out, zeroes



def crawl(data: Dict[Point, int], start: Point) -> int:
    nines: Set[Point] = set()
    def inner_crawl(current_point: Point, seen: Set[Point]):
        if current_point not in data:
            return
        if data[current_point] == 9:
            nines.add(current_point)
            return

        seen.add(current_point)
        # else we need to look in all three directions .. 
        left = Point(current_point.row, current_point.col - 1)
        right = Point(current_point.row, current_point.col + 1)
        up = Point(current_point.row - 1, current_point.col)
        down = Point(current_point.row + 1, current_point.col)


        if left not in seen and left in data and (data[left] - data[current_point] == 1):
            inner_crawl(left, seen)
        if right not in seen and right in data and (data[right] - data[current_point]) == 1:
            inner_crawl(right, seen)
        if up not in seen and up in data and (data[up] - data[current_point]) == 1:
            inner_crawl(up, seen)
        if down not in seen and down in data and (data[down] - data[current_point]) == 1:
            inner_crawl(down, seen)

    inner_crawl(start, set())
    return len(nines)



def path_to_str(path: List[Point]) -> str:
    out = ""
    for point in path:
        out += f'({point.row}-{point.col})'
    return out


def crawl_paths(data: Dict[Point, int], start: Point) -> int:
    nines_paths: List[str] = []
    def inner_crawl(current_point: Point, seen: List[Point]):
        if current_point not in data:
            return
        if data[current_point] == 9:
            nines_paths.append(path_to_str(seen))
            return

        seen = copy.deepcopy(seen)
        seen.append(current_point)
        # else we need to look in all three directions .. 
        left = Point(current_point.row, current_point.col - 1)
        right = Point(current_point.row, current_point.col + 1)
        up = Point(current_point.row - 1, current_point.col)
        down = Point(current_point.row + 1, current_point.col)


        if left not in seen and left in data and (data[left] - data[current_point] == 1):
            inner_crawl(left, seen)
        if right not in seen and right in data and (data[right] - data[current_point]) == 1:
            inner_crawl(right, seen)
        if up not in seen and up in data and (data[up] - data[current_point]) == 1:
            inner_crawl(up, seen)
        if down not in seen and down in data and (data[down] - data[current_point]) == 1:
            inner_crawl(down, seen)

    inner_crawl(start, [])
    return len(nines_paths)

def solve1(data: Dict[Point,int], zeroes: Set[Point]) -> int:
    result = 0

    for zero in zeroes:
        nines = crawl(data, zero)
        result += nines

    return result


def solve2(data: Dict[Point, int], zeroes: Set[Point]) -> int:
    result = 0

    for zero in zeroes:
        nines = crawl_paths(data, zero)
        result += nines

    return result

if __name__ == '__main__':
    data, zeroes = get_input()
    print(solve1(data,zeroes))
    print(solve2(data,zeroes))
