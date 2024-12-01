import functools

def get_input(): 
    lines = open('input.txt').read().split('\n')
    l1, l2 = [], []
    for line in lines: 
        if line == "":
            continue
        parts = line.split("  ") # input file has two spaces lol 
        l1.append(int(parts[0]))
        l2.append(int(parts[1]))
    return (l1, l2)

def count_sorted(needle: int, lst: [int]) -> int:
    count = 0
    for n in lst:
        if n > needle:
            return count
        if n == needle:
            count += 1
    return count

def solve2(l1, l2) -> int:
    l2 = sorted(l2)
    return sum(map(lambda k: k * count_sorted(k, l2), l1))


def solve1(l1, l2) -> int:
    zipped = list(zip(sorted(l1), sorted(l2)))
    return functools.reduce(lambda a, b: a + (abs(b[0] - b[1])), zipped, 0)

if __name__ == '__main__':
    l1, l2 = get_input()
    print(solve1(l1, l2))
    print(solve2(l1, l2))


