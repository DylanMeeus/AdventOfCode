



def getnumbers(string):
    nums = []
    if ".." in string:
        rangeps = string.split("..")
        for i in range(int(rangeps[0]), int(rangeps[1])+1):
            nums.append(i)
    else:
        nums.append(int(string))
    return nums


def parse(data):
    world = {}
    world[(500,0)] = "+"
    for line in data:
        if line == "":
            continue
        parts = line.split(",")
        fst = getnumbers(parts[0][2:])
        snd = getnumbers(parts[1][2:])
        rows = fst if parts[0][0] == "y" else snd
        columns = fst if parts[0][0] == "x" else snd
        for y in rows:
            for x in columns:
                world[(x,y)] = "#"
    return world

def draw(world):
    # find the boundary to draw it :
    minx = min(list(map(lambda k: k[0], world.keys())))
    maxx = max(list(map(lambda k: k[0], world.keys())))
    miny = min(list(map(lambda k: k[1], world.keys())))
    maxy = max(list(map(lambda k: k[1], world.keys())))
    out = ""
    for y in range(miny-1, maxy+1):
        for x in range(minx-1, maxx+1):
            if (x,y) in world.keys():
                out += world[(x,y)]
            else:
                out += "."
        out += "\n"
    print(out)


if __name__ == '__main__':
    data = open("test.txt","r").read().split("\n")
    w = parse(data)
    draw(w)
