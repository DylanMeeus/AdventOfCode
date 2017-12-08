# Small DSL to parse AOC8



# define operations, and equality comparisons
# but first create a register



varvalues = {}

def scan_variables(data):
    variables = list(map(lambda k: k.split(' ')[0],data))
    for variable in variables:
        varvalues[variable] = 0

def getInput():
    f = open("input8.txt",'r')
    data = f.read()
    return data.split('\n')[:-1]


def solve():
    data = getInput()
    scan_variables(data)
    # parse instructions with the variables
    for line in data:
        segments = line.split(' if ')
        operation = segments[0]
        predicate = segments[1]
        p = parse_predicate(predicate)
        if p:
            parse_operation(operation)

    print(max(varvalues.values()))


def resolve_variable(variable):
    if variable in varvalues.keys():
        return varvalues[variable]
    return variable


# Parser for operations
def parse_operation(operation):
    parts = operation.split(' ')
    variable = parts[0]
    op = parts[1]
    value = int(parts[2])
    if op == "inc":
        varvalues[variable] = varvalues[variable] + value
    else:
        varvalues[variable] = varvalues[variable] - value
    

# Parser for predicates
def parse_predicate(predicate):
    parts = predicate.split(' ')
    lhs = int(resolve_variable(parts[0]))
    eqOperator = parts[1]
    rhs = int(resolve_variable(parts[2]))
    return compare_ops[eqOperator](lhs,rhs)



def eq(lhs, rhs):
    return lhs == rhs

def neq(lhs, rhs):
    return lhs != rhs

def sm(lhs, rhs):
    return lhs < rhs

def gr(lhs, rhs):
    return lhs > rhs

def sme(lhs, rhs):
    return lhs <= rhs

def gre(lhs,rhs):
    return lhs >= rhs

compare_ops = {
        "==":eq,
        "!=":neq,
        "<":sm,
        ">":gr,
        "<=":sme,
        ">=":gre
        }
solve()
