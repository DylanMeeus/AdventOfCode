import re

def find_cycle(pattern, graph, start_node):
    cycle_counter = 0
    pattern_ptr = 0
    found_cycle = False
    current = start_node
    encounters = {}


    z_indices = []

    while not found_cycle:
        current = graph[current][0] if pattern[pattern_ptr] == 'L' else graph[current][1]
        cycle_counter += 1

        # determine if we have a cycle
        tpl = (current, pattern_ptr)

        if current.endswith('Z'):
            z_indices.append(cycle_counter)

        if tpl in encounters:
            found_cycle = True
        else:
            encounters[tpl] = True

        pattern_ptr = pattern_ptr + 1 if pattern_ptr < len(pattern) - 1 else 0


    return (cycle_counter, z_indices)
    

def solve2(pattern, graph):
    current_nodes = list(filter(lambda k: k.endswith('A'), graph.keys()))


    z_indices = []
    for node in current_nodes:
        z_indices.append(find_cycle(pattern, graph, node))


    print(z_indices)


    i = 15529 
    while True:
        all_z = True
        for idx in z_indices:
            if all_z == False:
                continue
            result_idx = i % (idx[0] + 1 )
            #print(f'i = {i} and result_idx = {result_idx} for {idx[0]}')
            if result_idx not in idx[1]: # we did not hit Z
                all_z = False
        if all_z:
            return i
        i += 15531



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
    lines = open('input.txt').read().split('\n')
    inputs = parse(lines)
    #print(solve1(inputs[0], inputs[1]))
    print(solve2(inputs[0], inputs[1]))
