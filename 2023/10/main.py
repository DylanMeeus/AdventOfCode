from enum import Enum 

class direction(Enum):
    UP = 1,
    DOWN = 2,
    LEFT = 3,
    RIGHT = 4




def parse(lines):
    graph = {}
    for row, line in enumerate(lines):
        for col, char in enumerate(line):
            graph[(row,col)] = char
    return graph



def solve1(graph):
    start = None
    for key, value in graph.items():
        if value == "S":
            start = key

    move_right = lambda k: (k[0], k[1] + 1)
    move_left = lambda k: (k[0], k[1] - 1)
    move_up = lambda k : (k[0] - 1, k[1])
    move_down = lambda k : (k[0] + 1, k[1])

    move = {
            direction.DOWN: move_down,
            direction.UP: move_up,
            direction.LEFT : move_left,
            direction.RIGHT : move_right
    }


    move_modification = {

            # north - east 
            ("L", direction.DOWN): direction.RIGHT,
            ("L", direction.LEFT): direction.UP,

            # north - west
            ("J", direction.RIGHT): direction.UP,
            ("J", direction.DOWN): direction.LEFT,

            # south - west
            ("7", direction.RIGHT): direction.DOWN,
            ("7", direction.UP): direction.LEFT,

            # south - east 
            ("F", direction.LEFT): direction.DOWN,
            ("F", direction.UP): direction.RIGHT,

    }

    # quick visual inspection of the text file shows us we can only go right or left at the start
    d = direction.RIGHT
    steps = 1
    current_node = move[d](start)
    while current_node != start:
        current_node = move[d](current_node)
        # should we change direction? 
        steps += 1

        if current_node == start:
            break

        symbol = graph[current_node]
        if symbol == "-" or symbol == "|":
            # no change in direction
            continue

        d = move_modification[(symbol, d)]

    # half-way from each end
    return steps // 2
        




if __name__ == '__main__':
    lines = open('input.txt').read().split('\n')[:-1]
    graph = parse(lines)
    print(solve1(graph))
