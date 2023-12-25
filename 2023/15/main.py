from functools import reduce



def get_box(part):
    chars = [x for x in filter(lambda k: k != "\n", part)]
    out = list(map(lambda k: ord(k), chars))
    val = reduce(lambda x,y: ((x + y) * 17) % 256, out, 0)
    return val


def solve1(parts):
    _sum = 0
    for part in parts:
        chars = [x for x in filter(lambda k: k != "\n", part)]
        out = list(map(lambda k: ord(k), chars))
        val = reduce(lambda x,y: ((x + y) * 17) % 256, out, 0)
        _sum += val

    return _sum



def solve2(parts):
    for part in parts:
        box = get_box(part)
        print(box)

if __name__ == '__main__':
    parts = open('test_input.txt').read().split(',')
    #print(solve1(parts))
    print(solve2(parts))
