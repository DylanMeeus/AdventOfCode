# Solution for day 11



# Map it to a Cartesian System
# Adjust for floating coordinates by *2

def getInput():
    f = open("input11.txt",'r')
    inp = f.read()[:-2:]
    return inp.split(",")

def solve():
    data = getInput()
    directions = {"nw":(-1,1),"n":(0,2),"ne":(1,1),"se":(1,-1),"s":(0,-2),"sw":(-1,-1)}

    pos = (0,0)
    for d in data:
        double = directions[d]
        pos = (pos[0]+double[0],pos[1]+double[1])
    
    # Find Manhattan Distance!
    distance = (abs(pos[0]) + abs(pos[1])) / 2
    print(distance)



def solve2():
    data = getInput()
    directions = {"nw":(-1,1),"n":(0,2),"ne":(1,1),"se":(1,-1),"s":(0,-2),"sw":(-1,-1)}

    pos = (0,0)
    furthest = 0
    for d in data:
        double = directions[d]
        pos = (pos[0]+double[0],pos[1]+double[1])
        distance = (abs(pos[0]) + abs(pos[1])) / 2
        if distance > furthest:
            furthest = distance
    
    # Find Manhattan Distance!
    print(furthest)

solve()
solve2()
