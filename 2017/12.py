# Solution to Day 12


class group:
    def __init__(self, g):
        self.group = g

    def contains(self, x):
        return x in self.group

    def add(self, xs):
        self.group.extend(xs)

def getInput():
    f = open('input12.txt','r')
    data = f.read()
    return data.split("\n")[:-1]


def solve2():
    data = getInput()
    heads = []
    
    for d in data:
        head = d.split(' ')[0]
        heads.append(head)
        heads = list(map(lambda k: int(k),heads))
    
    groups = []
    while len(heads) > 0:
        # Create a chain
        chain = create_link(heads[0],[],data)
        print(chain)
        if chain == None:
            continue
        # filter unique elements
        chain = map(lambda k: int(k),chain)
        unique = set(chain)
        for existing_group in groups:
            for el in unique:
                if existing_group.contains(el):
                    existing_group.add(unique)
                    print("going to break")
                    break
        else:
            gr = group(unique)
            groups.append(gr)
        print(len(heads))
        heads = list(filter(lambda k: k not in unique,heads))
        print(len(heads))
    print(len(groups))

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


#solve()
solve2()
