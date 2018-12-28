


def parse():
    data = open("input.txt","r").read().split("\n")
    vecs = []
    for line in data:
        if line == "":
            continue
        parts = line.split("x")
        vecs.append((int(parts[0]), int(parts[1]), int(parts[2])))
    return vecs


def solve(data):
    s = 0
    for entry in data:
        l,w,h = entry
        a = (2*l*w) + (2*w*h) + (2*h*l)
        # multiply all sides, divide by largest. result is multiplication of two smallest :)
        a += (l*w*h)/max(entry)
        s += a
    return s

def solve2(data):
    # how much feet of ribbon do they need
    s = 0
    for entry in data:
        l,w,h = entry
        # find the two smallest faces
        s += (l*2 + w*2 + h*2) - (max(entry)*2)
        s += (l*w*h)
    return s

if __name__ == '__main__':
    data = parse()
    print(solve(data))
    print(solve2(data))
    # 2*l*w + 2*w*h + 2*h*l





