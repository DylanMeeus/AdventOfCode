

def validate_page(rule, seen, all_pages) -> bool:
    for predecesor in rule: 
        if predecesor not in seen and predecesor in all_pages:
            return False
    return True


""" solve the first puzzle based on the rules and pages """
def solve1(rules, pages) -> int:
    result = 0 

    for page_set in pages:
        valid_page_set = True
        seen = set()
        for page in page_set:
            if page in rules:
                applicable_rule = rules[page]
                if not validate_page(applicable_rule, seen, page_set):
                    valid_page_set = False
            seen.add(page)
        if valid_page_set:
            mid = len(page_set) // 2
            result += page_set[mid]


    return result


def no_edge(S, G) -> int:
    for value in S:
        if not G in 


def fix_page_set(page_set, rules) -> [int]:
    # construct the sub-graph first
    G = {} 
    
    nub_values = set()

    for page in page_set:
        nub_values.add(page)
        if page in rules:
            G[page] = []
            for value in rules[page]:
                if value in page_set:
                    G[page].append(value)

    # now build the page_set in reverse against G, by finding the node with no "After" (edge)
    # then delete this node from all rules.. 


    return page_set


def solve2(rules, pages) -> int:
    """ build a graph from the 'end' node, and delete each consumed node.. until N/2 nodes are deleted """ 
    result = 0

    for page_set in pages:
        valid_page_set = True
        seen = set()
        for page in page_set:
            if page in rules:
                applicable_rule = rules[page]
                if not validate_page(applicable_rule, seen, page_set):
                    valid_page_set = False
            seen.add(page)
        if not valid_page_set:
            page_set = fix_page_set(page_set, rules)
            mid = len(page_set) // 2
            result += page_set[mid]



    return result


def get_input():
    # parse out the input in 'rules' and 'actual input'
    lines = open('test_input.txt').read().split('\n')

    pages = []
    # rules encodes a key (after) a value (before)
    # this allows us to do a O(n) solution for solve1
    rules = {}

    for line in lines:
        if "|" in line:
            parts = list(map(lambda k: int(k), line.split("|")))
            if parts[1] not in rules:
                rules[parts[1]] = [parts[0]]
                continue
            rules[parts[1]].append(parts[0])

        elif "," in line:
            pages.append(list(map(lambda k: int(k), line.split(","))))
        else:
            # invalid line
            None
    return rules, pages




if __name__ == '__main__':
    r, p = get_input()
    print(r)
    print(solve1(r,p))
    print(solve2(r,p))

