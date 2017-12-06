# Advent of code day 6









def getInput():
    f = open("input6.txt","r")
    strIn = f.read()
    strVal = strIn.split("\t")
    print(strVal)
    return list(map(lambda k: int(k),strVal))



def largestIndex(data):
    li = 0
    for i in range(len(data)):    
        if data[i] > data[li]:
            li = i
    return li

def redistribute(registers, i):
    reglen = len(registers)
    capture = registers[i]
    registers[i] = 0
    i = (i + 1) % reglen
    while capture > 0:
        registers[i] = registers[i] + 1
        i = (i + 1) % reglen
        capture -= 1
    return registers


def solve():
    registers = (getInput())
    store = []
    store.append(registers)
    foundMatch = False
    cycles = 0 
    while not foundMatch:
        ind = largestIndex(registers)
        registers = redistribute(list(registers), ind) # new list to create a copy
        cycles += 1
        if registers in store:
            foundMatch = True
        store.append(registers)
    print(cycles)
    







solve()
