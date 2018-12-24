from functools import reduce

depth = 510 
target = (9,731)


def solve():
    # create the map
    # Need to make a map of (9,731)
    world = {}
    for y in range(10+1): # test
        for x in range(10+1): # test
            world[(x,y)] = get_erosion(geo_index((x,y),world))
    #print(world)
    # calculate the risk
    risk = reduce(lambda y,x: y + (x % 3), world.values())
    print(risk)
    print(stdout(world))
     


def stdout(world):
    out = ""
    for y in range(11):
        for x in range(11):
            val = world[(x,y)] % 3
            if val == 0:
                out += "."
            elif val == 1:
                out += "="
            elif val == 2:
                out += "|"
        out += "\n"
    return out

def geo_index(point, world):
    x,y = point[0], point[1]
    if x == 0 and y == 0:
        return 0
    if y == 0:
        return x * 16807
    if x == 0:
        return y * 48271
    # else
    left = (x - 1, y)
    up = (x,y - 1)
    leftvalue = geo_index(left, world)
    upvalue = geo_index(up, world)
    return get_erosion(leftvalue) * get_erosion(upvalue)

def get_erosion(geodex):
    return (geodex + depth) % 20183 



if __name__ == '__main__':
    solve()
