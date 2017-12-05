# Problem 5 of advent of code


# find out how many steps it take to escape a list
# "0 3 0 1 -3"


def getInput():
    f = open("input5.txt",'r')
    inp = f.read()
    strNumbers = inp.split("\n")
    return list(map(lambda k: int(k), strNumbers[:-1]))

def solve():
    print("solving problem")
    input = getInput()
    i = 0
    steps = 0
    while(i < len(input)):
        jmp = input[i]
        input[i] = jmp + 1
        if jmp >= 3:            # Check for part2 of day5. Run without check for part 1
            input[i] = jmp - 1 
        
        i += jmp
        steps += 1

    print(steps)

solve()
