from dataclasses import dataclass


@dataclass
class Point:
    row: int
    col: int

    def __hash__(self):
        return hash((self.row, self.col))

def to_map(lines: [str]):
    out = {}
    for row in range(0, len(lines)):
        for col in range(0, len(lines[row])):
            p = Point(row,col)
            out[p] = lines[row][col]

    return out


def solve1(point_cloud):
    count = 0

    # filter for all the Xs, then scan in each direction to see if it makes up "XMAS"
    for point, value in point_cloud.items():
        if value == "X":
            for func in get_point_funcs():
                new_points = func(point)
                word = get_string(point_cloud, new_points)
                if word == "XMAS":
                    count += 1

    return count


def safe_get(m, p) -> str:
    if p in m:
        return m[p]
    return "_"

def solve2(point_cloud):
    # filter for 'A'
    count = 0
    for point, value in point_cloud.items():
        if value == "A":
            top_left = Point(point.row-1, point.col-1)
            top_right = Point(point.row - 1, point.col + 1)
            bottom_left = Point(point.row + 1, point.col - 1)
            bottom_right = Point(point.row + 1, point.col + 1)
            f = lambda k: safe_get(point_cloud, k)

        
        ###
        # M  M    S  S   M  S   S  M
        # S  S    M  M   M  S   S  M
            if f(top_left) == 'M' and f(top_right) == 'M' and f(bottom_left) == 'S' and f(bottom_right) == 'S':
                count += 1

            if f(top_left) == 'S' and f(top_right) == 'S' and f(bottom_left) == 'M' and f(bottom_right) == 'M':
                count += 1

            if f(top_left) == 'M' and f(top_right) == 'S' and f(bottom_left) == 'M' and f(bottom_right) == 'S':
                count += 1

            if f(top_left) == 'S' and f(top_right) == 'M' and f(bottom_left) == 'S' and f(bottom_right) == 'M':
                count += 1

    return count





def get_string(point_cloud, points) -> str:
    out = "" 
    for point in points:
        if point in point_cloud:
            out += point_cloud[point]
    return out



def get_point_funcs():
    return [ get_points_up, 
            get_points_down,
            get_points_left,
            get_points_right,
            get_points_left_up,
            get_points_left_down,
            get_points_right_up,
            get_points_right_down]

def get_points_up(p): 
    # return all points based on this...
    return [p,  # X 
            Point(p.row - 1, p.col), # M 
            Point(p.row - 2, p.col), # A
            Point(p.row - 3, p.col) #  S
    ]

def get_points_down(p): 
    # return all points based on this...
    return [p,  # X 
            Point(p.row + 1, p.col), # M 
            Point(p.row + 2, p.col), # A
            Point(p.row + 3, p.col) #  S
    ]

def get_points_left(p): 
    # return all points based on this...
    return [p,  # X 
            Point(p.row , p.col - 1), # M 
            Point(p.row , p.col - 2), # A
            Point(p.row , p.col - 3) #  S
    ]


def get_points_right(p): 
    # return all points based on this...
    return [p,  # X 
            Point(p.row , p.col + 1), # M 
            Point(p.row , p.col + 2), # A
            Point(p.row , p.col + 3) #  S
    ]


def get_points_left_up(p):
    return [p,  # X 
            Point(p.row - 1 , p.col - 1), # M 
            Point(p.row - 2 , p.col - 2), # A
            Point(p.row - 3, p.col - 3) #  S
    ]


def get_points_right_up(p):
    return [p,  # X 
            Point(p.row - 1 , p.col + 1), # M 
            Point(p.row - 2 , p.col + 2), # A
            Point(p.row - 3, p.col + 3) #  S
    ]

def get_points_left_down(p):
    return [p,  # X 
            Point(p.row + 1 , p.col - 1), # M 
            Point(p.row + 2 , p.col - 2), # A
            Point(p.row + 3, p.col - 3) #  S
    ]

def get_points_right_down(p):
    return [p,  # X 
            Point(p.row + 1, p.col + 1), # M 
            Point(p.row + 2, p.col + 2), # A
            Point(p.row + 3, p.col + 3) #  S
    ]
def get_input() -> [str]: 
    lines = open('input.txt').read().split('\n')
    return lines

if __name__ == '__main__':
    mapped = to_map(get_input())
    print(solve1(mapped))
    print(solve2(mapped))
