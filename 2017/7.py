# Solution to question 7



def getInput():
    f = open("input7.txt",'r')
    lines = f.read()
    return lines.split("\n")


def solve():
    data = getInput()[:-1]
    # For each 'node' with '->', we check if the start appears in any other 'tail'
    data = list(filter(lambda k: '->' in k,data))
    heads = list(map(lambda k: k.split(' ')[0],data))
    tails = list(map(lambda k: k.split('->')[1],data))
    for head in heads:
        head_in_tail = False
        for tail in tails:
            if head in tail:
                head_in_tail = True 
                continue
        if not head_in_tail:
            print(head)
            exit()
            


solve()
