directions = {
        "N": (0,-1),
        "E": (1, 0),
        "S": (0, 1),
        "W": (-1,0)
    }

def add(fst, snd):
    x,y = fst[0] + snd[0], fst[1] + snd[1]
    return (x,y)

def solve2(rgx):
    return len(list(filter(lambda k: k >= 1000, solve(rgx).values())))

def solve(rgx): 
    stack = []
    current = (0,0)
    rooms = set() 
    groups = []
    distance = 0
    distances = {current: 0}
    for char in rgx:
        if char in "NESW":
            # change direction
            distance += 1
            current = add(current, directions[char])
            if current not in distances or distance < distances[current]:
                distances[current] = distance
            rooms.add(current)
        elif char == "(":
            groups.append((distance, current))
        elif char == "|":
            distance, current = groups[len(groups)-1]
        elif char == ")":
            groups.pop()
    return distances

if __name__ == '__main__':
    rgx = open("input.txt",'r').read()
    print(max(solve(rgx).values()))
    print(solve2(rgx))

