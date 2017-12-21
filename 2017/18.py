# Solution for day 18 of AoC



def getInput():
    f = open('input18.txt','r')
    return f.read().split("\n")[:-1]


def solve():
    data = getInput()
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

class Program:

    def __init__(self,pid, initial_data):
        self.pid = pid
        self.data = initial_data
        self.variables = {}
        for line in initial_data:
            variable = line.split(' ')[1]
            if variable in 'abcdefghijklmnopqrstuvwxyz':
                self.variables[variable] = 0
        self.line = 0
        self.terminated = False
        self.variables['p'] = pid
        print(self.variables)
        
        # Initialize commands 
        self.commands = {
                'set' : self.set_value,
                'add' : self.add,
                'mul' : self.mul,
                'mod' : self.mod,
                'rcv' : self.rcv,
                'snd' : self.snd
                }

        self.OTHER = None
        self.buffer = []
        self.on_hold = False
        self.send_values = 0

    def exec_instruction(self):
        if self.on_hold or self.terminated:
            return

        instruction = self.data[self.line]
        parts = instruction.split(" ")
        cmd = parts[0]
        # Process possible jump
        if cmd == "jgz":
            if self.resolve(self.variables, parts[1]) > 0:
                self.line = self.line + self.resolve(self.variables, parts[2])
                return
        # Process other instructions
        else:
            if len(parts) == 2:
                print(instruction)
                self.commands[cmd](parts[1])
            elif len(parts) == 3:
                self.commands[cmd](parts[1],self.resolve(self.variables,parts[2]))
        self.line += 1
        if self.line >= len(self.data):
            self.terminated = True


    # Instructions that mutate state
    def set_value(self, variable, value):
        self.variables[variable] = int(value)

    def add(self, variable, value):
        self.variables[variable] = self.variables[variable] + int(value)

    def mul(self,variable,value):
        self.variables[variable] = self.variables[variable] * int(value)

    def mod(self,variable,value):
        self.variables[variable] = int(self.variables[variable]) % int(value)

    def rcv(self,variable):
        if len(self.buffer) == 0:
            self.on_hold= True
            self.line = self.line - 1
        else:
            self.variables[variable] = self.buffer[0]
            self.buffer = self.buffer[1:]
            
    def snd(self,variable):
        self.send_values += 1
        self.OTHER.ping(self.resolve(self.variables,variable))


    def resolve(self,variables, possible_variable):
        if possible_variable in 'abcdefghijklmnopqrstuvwxyz':
            return int(variables[possible_variable])
        return int(possible_variable)

    def ping(self,value):
        self.buffer.append(value)
        self.on_hold = False

def solve2():
    data = getInput()
    one = Program(0,data)
    two = Program(1,data)
    one.OTHER = two
    two.OTHER = one

   
    while (not one.terminated) and (not two.terminated):
        # Deadlock check
        if one.on_hold and two.on_hold:
            print("DEADLOCKED")
            print(two.send_values)
            exit()
        one.exec_instruction()
        two.exec_instruction()
solve2()
