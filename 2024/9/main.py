



def get_input() -> str:
    return open('input.txt').read()


def compress(s: str) -> str:
    """ compress using a two-pointer approach """ 
    start = 0
    end = len(s) - 1

    l = list(s)
    while start < end: 
        if l[start] != '.':
            start += 1
            continue
        else:
            if l[end] == '.':
                end -= 1
                continue
            else:
                # lwap them.. 
                copy = l[start]
                l[start] = l[end]
                l[end] = copy
                start += 1

    return ''.join(l)

def expand(s: str) -> str:
    out = ''
    idx = 0
    for i, freq in enumerate(s):
        if not freq.isdigit():
            continue
        if i % 2 == 0:
            out += str(idx) * int(freq)
            idx += 1
        else:
            out += "." * int(freq)
    return out
        

def checksum(data: str) -> int:
    out = 0
    for idx, file_id in enumerate(data):
        if file_id == '.':
            return out
        out += (idx * int(file_id))
    return out

def solve1(data) -> int:
    expanded = expand(data)
    compressed = compress(expanded)
    return checksum(compressed)


if __name__ == '__main__':
    data = get_input()
    print(solve1(data))
