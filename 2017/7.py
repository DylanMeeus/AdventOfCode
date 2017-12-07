# Solution to question 7




class node:
    def __init__(self,name,weight, parent = None):
        self.name = name.strip()
        self.weight = weight
        self.children = []
        self.parent = parent
        
    def total_weight(self):
        w = self.weight
        for child in self.children:
            w += child.total_weight()
        return w


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
   
    print(r.total_weight())
    find_wrong_weight(r)
    exit()


def find_wrong_weight(n):
    if len(n.children) == 0:
        return
    mx = 0
    mxn = n
    for child in n.children:
        ct = child.total_weight()
        print(n.name + " :: " + child.name + ": " + str(ct))
        if ct > mx:
            mx = ct
            mxn = child
    print("=================")
    find_wrong_weight(mxn)

def dedup(n,p):
    p.append(n)
    for child in n.children:
        dedup(child,p)
    return p


def build_tree(n,nodes,axons):
    for axon in axons:
        if axon.startswith(n.name):
            ps = axon.split('->')[1].split(',')
            for p in ps:
                pn = get_node(p,nodes)
                pn = build_tree(pn,nodes,axons)
                n.children.append(pn)
                pn.parent=n

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
