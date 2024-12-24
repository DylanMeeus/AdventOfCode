

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




def get_input():
    # parse out the input in 'rules' and 'actual input'
    lines = open('input.txt').read().split('\n')

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

