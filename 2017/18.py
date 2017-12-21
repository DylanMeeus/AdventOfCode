# Solution for day 18 of AoC



def getInput():
    f = open('input18.txt','r')
    return f.read().split("\n")[:-1]


# store the last played sound
played_sounds = {}

frequency = -1
# Instructions that mutate state
def set_value(variables, variable, value):
    variables[variable] = int(value)

def add(variables, variable, value):
    variables[variable] = variables[variable] + int(value)

def mul(variables,variable,value):
    variables[variable] = variables[variable] * int(value)

def mod(variables,variable,value):
    variables[variable] = int(variables[variable]) % int(value)

def rcv(variables,variable):
    if variables[variable] == 0:
        return
    if frequency > 0:
        print(frequency)
        exit()
    variables[variable] = frequency
        
def snd(variables,variable):
    global frequency
    frequency = int(variables[variable])
    played_sounds[variable] = int(variables[variable])

# End instructions that mutate state

def resolve(variables, possible_variable):
    if possible_variable in 'abcdefghijklmnopqrstuvwxyz':
        return int(variables[possible_variable])
    return int(possible_variable)

def solve():
    data = getInput()
    
    commands = {
            'set' : set_value,
            'add' : add,
            'mul' : mul,
            'mod' : mod,
            'rcv' : rcv,
            'snd' : snd
            }


    variables = {}
    # Create the variable map the first sweep
    for line in data:
        variable = line.split(' ')[1]
        if variable in 'abcdefghijklmnopqrstuvwxyz':
            variables[variable] = 0 

    line = 0
    while line < (len(data)):
        instruction = data[line]
        parts = instruction.split(" ")
        cmd = parts[0]
        # Process possible jump
        if cmd == "jgz":
            if resolve(variables, parts[1]) > 0:
                line = line + resolve(variables, parts[2])
                continue
        # Process other instructions
        else:
            if len(parts) == 2:
                print(instruction)
                commands[cmd](variables,parts[1])
            elif len(parts) == 3:
                commands[cmd](variables,parts[1],resolve(variables,parts[2]))
        line += 1


solve()
