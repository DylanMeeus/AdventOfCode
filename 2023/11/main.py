import numpy as np
import math




def universe_to_string(universe):
    output = ""
    for row, entry in enumerate(universe):
        output += "".join(entry) + "\n"
    return output




def expand_universe(universe):
    # find all rows and columsn to be expanded
    rows_to_expand = []
    for row, line in enumerate(universe):
        if line.count(".") == len(line):
            rows_to_expand.append(row)


    # find all rows and columsn to be expanded
    universe_T = transpose(universe)


    cols_to_expand = []
    for row, line in enumerate(universe_T):
        if line.count(".") == len(line):
            cols_to_expand.append(row)


    print(cols_to_expand)

    output = ""
    for row, content in enumerate(universe):
        line_out = ""
        for col, char in enumerate(content):
            line_out += char
            if col in cols_to_expand:
                line_out += char

        output += line_out + "\n"
        if row in rows_to_expand:
            output += line_out + "\n"

    return (output, rows_to_expand, cols_to_expand)




def parse_to_array(lines):
    universe = []
    for row, line in enumerate(lines):
        entry = []
        for col, char in enumerate(line):
            entry.append(char)
        universe.append(entry)
    return universe


def transpose(matrix):
    np_arr = np.array(matrix)
    return [ [matrix[j][i] for j in range(len(matrix))] for i in range(len(matrix[0])) ]


def to_graph(lines):
    graph = {}
    for row, line in enumerate(lines):
        for col, char in enumerate(line):
            if char == "#":
                graph[(row,col)] = "#"
    return graph


def solve2(graph, rows, cols):
    _sum = 0
    for star, _ in graph.items():
        for other_star, _ in graph.items():
            if star == other_star:
                continue

            ## find the distance
            deltas = ((other_star[0] - star[0]) ** 2) + ((other_star[1] - star[1]) ** 2)
            dist = math.sqrt( deltas ) 

            # find the actual distance 

            if not (star[0] == other_star[0] or star[1] == other_star[1]):
                dist = abs(star[0] - other_star[0]) + abs(star[1] - other_star[1])

            # find out how many expanded rows and cols we crossed

            _count = 0
            for expanded_row in rows:
                if between(expanded_row, star[0], other_star[0]):
                    _count += 1

            for expanded_col in cols:
                if between(expanded_col, star[1], other_star[1]):
                    _count += 1

            _sum += (dist + (_count * 1_000_000)) - (_count)
    return _sum // 2


def between(x, start, end):
    if (x > start and x < end) or (x > end and x < start):
        return True

def solve1(graph):
    _sum = 0
    for star, _ in graph.items():
        for other_star, _ in graph.items():
            if star == other_star:
                continue

            ## find the distance
            deltas = ((other_star[0] - star[0]) ** 2) + ((other_star[1] - star[1]) ** 2)
            dist = math.sqrt( deltas ) 

            # find the actual distance 

            if not (star[0] == other_star[0] or star[1] == other_star[1]):
                dist = abs(star[0] - other_star[0]) + abs(star[1] - other_star[1])

            _sum += dist
    return _sum // 2



if __name__ == '__main__':
    lines = open('input.txt').read().split('\n')[:-1]
    U = parse_to_array(lines)
    expanded, rows, cols = expand_universe(U)
    expanded = expanded.split("\n")
    print(solve1(to_graph(U)))
    print(solve2(to_graph(U), rows, cols))

    
