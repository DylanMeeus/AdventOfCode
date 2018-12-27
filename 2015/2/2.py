


def parse():
    data = open("input.txt","r").read().split("\n")
    vecs = []
    for line in data:
        if line == "":
            continue
        parts = line.split("x")
        vecs.append((int(parts[0]), int(parts[1]), int(parts[2])))
    return vecs


if __name__ == '__main__':
    data = parse()
    # 2*l*w + 2*w*h + 2*h*l
    s = 0
    for entry in data:
        l,w,h = entry
        a = (2*l*w) + (2*w*h) + (2*h*l)
        # multiply all sides, divide by largest. result is multiplication of two smallest :)
        a += (l*w*h)/max(entry)
        s += a
    print(s)





