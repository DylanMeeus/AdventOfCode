




def derive_crc(code):
    parts = code.split(".")
    nums = list(filter(lambda k: k != 0, map(lambda k: len(k), parts)))
    return nums


def crc_eq(xs, ys):
    if len(xs) != len(ys):
        return False
    for i,v in enumerate(xs):
        if v != ys[i]:
            return False
    return True


def count_ways(line, crc):
    if "?" not in line:
        if crc_eq(derive_crc(line), crc):
            return 1
        else:
            return 0

    # else we have to create the CRC


    q = line.find('?')

    opt_a = line.replace('?', '.', 1)
    opt_b = line.replace('?', '#', 1)

    return count_ways(opt_a, crc) + count_ways(opt_b, crc)

    
        


def solve1(lines):
    _sum = 0
    for line in lines:
        parts = line.split(' ')

        code = parts[0]
        crc = list(map(lambda k: int(k), parts[1].split(",")))

        _sum += count_ways(code,crc)

    return _sum



if __name__ == '__main__':
    lines = open('input.txt').read().split('\n')[:-1]
    print(solve1(lines))

