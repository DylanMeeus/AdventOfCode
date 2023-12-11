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



def solve2(graph):
    start = None
    for key, value in graph.items():
        if value == "S":
            start = key


    # quick visual inspection of the text file shows us we can only go right or left at the start
    d = direction.RIGHT
    steps = 1
    current_node = move[d](start)

    collected = [current_node, start]

    while current_node != start:
        current_node = move[d](current_node)
        collected.append(current_node)
        # should we change direction? 
        steps += 1

        if current_node == start:
            break

        symbol = graph[current_node]
        if symbol == "-" or symbol == "|":
            # no change in direction
            continue

        d = move_modification[(symbol, d)]


    # cast a ray from each point, if ray intersect boundary % 2 == 0, outside. 

    _sum = 0
    for row in range(0, 140):
        for col in range(0,140):
            if (row, col) in collected:
                continue
            # otherwise cast a ray in any direction?? 

            dot = (row, col)
            count = 0
            while dot[0] < 140 and dot[1] < 140:
                if dot in collected and graph[dot] != "L" and graph[dot] != "7":
                    count += 1 
                dot = move_right(dot)
                dot = move_down(dot)


            if count % 2 == 1:
                _sum += 1


    return _sum



def solve1(graph):
    start = None
    for key, value in graph.items():
        if value == "S":
            start = key

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
    #print(solve1(graph))
    print(solve2(graph))



    # half-way from each end

    """ for debugging 
    string = ""
    for row in range(0, 140):
        for col in range(0,140):
            if (row, col) in collected:
                string += graph[(row,col)]
            else:
                string += "."
        string += "\n"

    print(string)
    """ 
