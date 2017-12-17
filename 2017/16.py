# solutions for day 16 of AoC
from functools import reduce
import copy

def getInput():
    f = open("input16.txt",'r')
    return f.read()[:-1].split(',')



def solve():
    method = {
                "s" : spin,
                "x" : exchange,
                "p" : partner
            }
    programs = list("abcdefghijklmnop")
    instructions = getInput()
    for instruction in instructions:
        print(instruction)
        programs = method[instruction[0]](instruction,programs)
    print(reduce(lambda a,b : str(a) + str(b),programs))

        


def solve2():
    method = {
                "s" : spin,
                "x" : exchange,
                "p" : partner
            }
    programs = list("abcdefghijklmnop")
    instructions = getInput()
    init_state = copy.deepcopy(programs)
    cycled = False
    cycles = 0
    while not cycled:
        for instruction in instructions:
            programs = method[instruction[0]](instruction,programs)
        cycles += 1
        if programs == init_state:
            cycled = True
            remaining_cycles = 1000000000 % cycles
            print(remaining_cycles)
            for i in range(remaining_cycles):
                for instruction in instructions:
                    programs = method[instruction[0]](instruction,programs)

            
    print(reduce(lambda a,b : str(a) + str(b),programs))



def spin(instruction, programs):
    amount = (int(instruction[1:]))
    tail = programs[-amount:]
    programs = (tail + programs)[:len(programs)]
    return programs

def exchange(instruction, programs):
    instruction = instruction[1:]
    parts = instruction.split('/')
    a = int(parts[0])
    b = int(parts[1])
    x = programs[a]
    programs[a] = programs[b]
    programs[b] = x
    return programs

    
    
    

def partner(instruction, programs):
    instruction = instruction[1:]
    parts = instruction.split('/')
    a = parts[0]
    b = parts[1]
    fst = programs.index(a)
    snd = programs.index(b)
    return exchange("x"+str(fst)+"/"+str(snd),programs)   



solve2()
