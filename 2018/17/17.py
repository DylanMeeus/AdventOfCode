



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
    world[(500,1)] = "|"
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

def add(tup1, tup2):
    return (tup1[0] + tup2[0], tup1[1] + tup2[1])

def solve(world):
    # keep expanding active water
    minx = min(list(map(lambda k: k[0], world.keys())))
    maxx = max(list(map(lambda k: k[0], world.keys())))
    miny = min(list(map(lambda k: k[1], world.keys())))
    maxy = max(list(map(lambda k: k[1], world.keys())))
    active = [(500,1)]
    up, down, left, right = (0,-1), (0,1), (-1,0), (1,0)
    while True:
        newactive = []
        # check if we still have valid active ones
        active = list(filter(lambda k: k[1] < maxy, active))
        active = list(set(active))
        if len(active) == 0:
            sump = 0
            for k in world.keys():
                if k[1] >= miny and k[1] <= maxy:
                    if world[k] in "~|":
                        sump+=1
            return sump

        for water in active:
            if add(water, down) not in world:
                world[add(water,down)] = "|"
                newactive.append(add(water,down))
            else:
                # water can't go down
                l = add(water, left)
                r = add(water, right)

                cbc = 0
                if l not in world:
                    world[l] = "|"
                    newactive.append(l)
                    cbc+=1
                if r not in world:
                    world[r] = "|"
                    newactive.append(r)
                    cbc+=1
                if cbc == 0 and isBounded(world, water):     
                    # select one up
                    y = water[1]
                    for x in range(water[0], maxx):
                        if (x,y-1) in world and world[(x,y-1)] == "|":
                            newactive.append((x,y-1))
                        if (x,y) in world and world[(x,y)] == "#":
                            break
                        else:
                            world[(x,y)] = "~"

                    x = water[0] 
                    while x >= minx:
                        if (x,y-1) in world and world[(x,y-1)] == "|":
                            newactive.append((x,y-1))
                        if (x,y) in world and world[(x,y)] == "#":
                            break
                        else:
                            world[(x,y)] = "~"
                        x -= 1


                            
        active = newactive


def isBounded(world, point):
    minx = min(list(map(lambda k: k[0], world.keys())))
    maxx = max(list(map(lambda k: k[0], world.keys())))
    miny = min(list(map(lambda k: k[1], world.keys())))
    maxy = max(list(map(lambda k: k[1], world.keys())))
    origx = point[0]
    y = point[1]
    rightbound = False
    leftbound = False
    for x in range(origx, maxx+1):
        if (x,y+1) not in world:
            rightbound = False
            break
        if (x,y) in world and world[(x,y)] == "#":
            rightbound = True
            break
    if not rightbound:
        return False
    x = origx
    while x >= minx:
        if (x,y+1) not in world:
            leftbound = False
            break
        if (x,y) in world and world[(x,y)] == "#":
            leftbound = True
            break
        x -= 1
    return rightbound and leftbound

    

    

if __name__ == '__main__':
    data = open(".txt","r").read().split("\n")
    w = parse(data)
    print(solve(w))
    #draw(w)

