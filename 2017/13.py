# Solution to day 13
import copy
class Scanner:
    
    def __init__(self):
        self.max_depth = 0
        self.depth = 1 # Yeah, it's not zero based. 
        self.increment = 1
    def move(self):
        self.depth += self.increment
        if self.depth == self.max_depth or self.depth == 1:
            self.increment *= -1


def getInput():
    x = "0: 3\n1: 2\n4: 4\n6: 4"
    #return x.split("\n")
    f = open("input13.txt",'r')
    return f.read().split("\n")[:-1]


def create_scanners(data):
    # find max, then make range, then fill range
    x = max(map(lambda k: int(k.split(':')[0]),data))
    scanners = []
    for i in range(x+1):
        scanners.append(Scanner())
    for line in data:
        parts = line.split(':')
        index = int(parts[0])
        depth = int(parts[1])
        scanners[index].max_depth = depth 
    return scanners

def solve():
    data = getInput()
    # Create scanners
    scanners = create_scanners(data)
    score = 0
    for start in range(len(scanners)):
        # find all the scanners 'start' moved on to, but exclude 0, also move the scanner
        print(str(start) + " scanner at: " + str(scanners[6].depth) + "::" +
            str(scanners[6].max_depth))
        if scanners[start].depth == 1:
            score += (start*scanners[start].max_depth)
        for scanner in scanners:
            scanner.move()
    print(score)



def solve2():
    data = getInput()
    # Create scanners
    scanners = create_scanners(data)
    solution = False

    delay = 0
    while not solution:
        # check if all positions are 1
        can_pass = True
        for pos in range(len(scanners)):
            if scanners[pos].max_depth == 0:
                scanners[pos].depth = -1
            else: 
                scanners[pos].depth = (delay + pos) % (2 * (scanners[pos].max_depth-1))

            if scanners[pos].depth == 0:
                can_pass = False

        if can_pass:
            print(delay)
            exit()
        delay += 1
    print(delay) 
#solve()
solve2()

