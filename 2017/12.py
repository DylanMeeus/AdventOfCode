# Solution to Day 12



def getInput():
    f = open('input12.txt','r')
    data = f.read()
    return data.split("\n")[:-1]


def solve():
    data = getInput()
    heads = []
    
    for d in data:
        head = d.split(' ')[0]
        heads.append(head)

    # Create a chain
    chain = create_link(0,[],data)
    # filter unique elements
    unique = set(chain)
    print(unique)
    print(len(unique))


processed_heads = []

def create_link(node, chain,data):
    if node in processed_heads:
        return 
    processed_heads.append(node)
    links = list(filter(lambda k: int(k.split(' ')[0]) == node,data))
    chain.append(node)
    for link in links:
        parts = link.split('<->')
        tails = parts[1].split(',')
        for tail in tails:
            create_link(int(tail),chain,data)

    return chain
solve()
