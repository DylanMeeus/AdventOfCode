from functools import reduce

depth = 11109 
target = (9,731)

#depth = 510
#target = (10,10)

geo_map = {}

def solve():
    # create the map
    # Need to make a map of (9,731)
    world = {}
    for y in range(target[1]+1): # test
        for x in range(target[0]+1): # test
            world[(x,y)] = get_erosion(geo_index((x,y)))
    # calculate the risk     
    risk = 0
    for v in world.values():
        risk += v % 3
    risk -= world[target] % 3
        
    print(risk)
    #print(stdout(world))
     


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

def geo_index(point):
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
    if left not in geo_map:
        geo_map[left] = geo_index(left)
    if up not in geo_map:
        geo_map[up] = geo_index(up)
    return get_erosion(geo_map[left]) * get_erosion(geo_map[up])

def get_erosion(geodex):
    return (geodex + depth) % 20183 



if __name__ == '__main__':
    solve()
