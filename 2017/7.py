# Solution to question 7




class node:
    def __init__(self,name,weight, parent = None):
        self.name = name.strip()
        self.weight = weight
        self.children = []
        self.parent = parent
        



def get_node(name,nodes):
    name = (name.strip())
    return list(filter(lambda k: k.name == name, nodes))[0]

def getInput():
    f = open("input7.txt",'r')
    lines = f.read()
    return lines.split("\n")



def solve2():
    """ find the node that needs to change to bring a balance to the force """
    data = getInput()[:-1]
    root_string = get_root()
    root_node = None
    nodes = []
    for line in data:
        parts = line.split(' ')
        n = (node(parts[0],int(parts[1][1:-1:])))
        nodes.append(n)
    root = get_node(root_string,nodes)
    axons = list(filter(lambda k: '->' in k, data))
    tree_root = build_tree(root,nodes,axons)
    r = tree_root
    # dedup them
    d = dedup(r,[])
    leafs = list(filter(lambda k: len(k.children) == 0,d))
    paths = []
    for leaf in leafs:
        parent = leaf.parent
        rpaths = [leaf]
        while parent != None:
            rpaths.append(parent)
            parent = parent.parent
        paths.append(rpaths)
    # find out which path is different at which point..
    testroot = list(filter(lambda k: k.parent == None, d))
    for p in paths:
        print(sum_weights(p))



def sum_weights(path):
    s = 0
    for n in path:
        s += n.weight
    return s

def dedup(n,p):
    p.append(n)
    for child in n.children:
        dedup(child,p)
    return p


def build_tree(n,nodes,axons):
    # check if n is an axon
    for axon in axons:
        if axon.startswith(n.name):
            parts = axon.split('->')[1].split(',')
            last = n
            if len(parts) > 0:
                for part in parts:
                    fn = get_node(part,nodes)
                    fn = build_tree(fn, nodes, axons)
                    fn.parent = last
                    last.children.append(fn)
                    last = fn

    return n

    
        
def get_root():
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
            return head
            
            

print(get_root())
solve2()
