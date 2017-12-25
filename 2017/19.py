# Solution to question 19.

# Return the [][]char 
def getInput():
    f = open('input19.txt','r')
    data = f.read()
    parts = data.split('\n')
    split_parts = list(map(lambda k: list(k), parts))
    return split_parts



def find_connection(data, location,original_move):
    left = (0,-1)
    right = (0,1)
    up = (-1,0)
    down = (1,0)
    positions = [left,right,up,down]
    reversed_move = (original_move[0] * -1, original_move[1] * -1)
    positions = list(filter(lambda k: k != reversed_move, positions))
    for position in positions:
        data_pos = (location[0] + position[0], location[1] + position[1])
        char = data[data_pos[0]][data_pos[1]]
        if char == '|' or char == '-':
            return position

    

def solve():
    data = getInput()

    row = 0
    column = 0
    first_row = data[0]
    for i in range(len(first_row)):
        if first_row[i] == '|':
            column = i
            break

    escaped = False
    move = (1,0)    # tuple indicating x (row) + y (column) change
    location = (0,column)
    chain = []
    while not escaped:
        location = (location[0] + move[0], location[1] + move[1])
        if location[0] > len(data):
            print("out of bounds")
            exit()
        if location[1] > len(data[location[0]]):
            print("out of other bounds")
            exit()
        
        char_at = data[location[0]][location[1]]
        if char_at not in ['|','-','+',' ']:
            chain.append(char_at)
            print(chain)

        
        if char_at == '+':
            # We have to change direction 
            move = find_connection(data,location,move)
            #print(move)
            
            


def solve2():
    data = getInput()

    row = 0
    column = 0
    first_row = data[0]
    for i in range(len(first_row)):
        if first_row[i] == '|':
            column = i
            break

    escaped = False
    move = (1,0)    # tuple indicating x (row) + y (column) change
    location = (0,column)
    chain = []
    total_steps = 1
    while not escaped:
        location = (location[0] + move[0], location[1] + move[1])
        
        char_at = data[location[0]][location[1]]
        if char_at not in ['|','-','+',' ']:
            chain.append(char_at)
            print(chain)




        total_steps += 1
        if char_at == 'P':
            print(total_steps)
            exit()
        if char_at == '+':
            # We have to change direction 
            move = find_connection(data,location,move)
            #print(move)
            
    

solve2()
