import re

def find_cycle(pattern, graph, start_node):
    cycle_counter = 0
    pattern_ptr = 0
    found_cycle = False
    current = start_node

    encounters = {}


    while not found_cycle:
        current = graph[current][0] if pattern[pattern_ptr] == 'L' else graph[current][1]
        cycle_counter += 1

        # determine if we have a cycle
        tpl = (current, pattern_ptr)

        if tpl in encounters:
            found_cycle = True
        else:
            encounters[tpl] = True

        pattern_ptr = pattern_ptr + 1 if pattern_ptr < len(pattern) - 1 else 0


    # with this cycle, find all Z indexes in this cycle

    print(current)
    z_indexes = 0 


    return cycle_counter
     

def solve2(pattern, graph):
    current_nodes = list(filter(lambda k: k.endswith('A'), graph.keys()))


    print(find_cycle(pattern, graph, current_nodes[0]))
    print(find_cycle(pattern, graph, current_nodes[1]))
    return 0

    step_counter = 0
    pattern_ptr = 0
    found_zzz = False
    while not found_zzz:

        # move all ghosts simultaneously
        for idx, current in enumerate(current_nodes):
            current = graph[current][0] if pattern[pattern_ptr] == 'L' else graph[current][1]
            current_nodes[idx] = current

        # check the state
        current_z_len = len(list(filter(lambda k: k.endswith('Z'), current_nodes)))
        if current_z_len == len(current_nodes):
            found_zzz = True
        pattern_ptr = pattern_ptr + 1 if pattern_ptr < len(pattern) - 1 else 0
        step_counter += 1
    return step_counter



def solve1(pattern, graph):
    step_counter = 0
    pattern_ptr = 0
    found_zzz = False
    current = 'AAA'
    while not found_zzz:
        current = graph[current][0] if pattern[pattern_ptr] == 'L' else graph[current][1]
        step_counter += 1
        if current == 'ZZZ':
            found_zzz = True
        pattern_ptr = pattern_ptr + 1 if pattern_ptr < len(pattern) - 1 else 0
    return step_counter




def parse(lines):
    pattern = lines[0]

    nav = {}

    for line in lines[1:]:
        if line == '':
            continue
        results = re.findall('\w+', line)
        nav[results[0]] = (results[1], results[2])
    return (pattern, nav)



if __name__ == '__main__':
    lines = open('test_input2.txt').read().split('\n')
    inputs = parse(lines)
    #print(solve1(inputs[0], inputs[1]))
    print(solve2(inputs[0], inputs[1]))
