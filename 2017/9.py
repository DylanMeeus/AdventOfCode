# solution to day 9 of advent of code





# Each time you close, {, the score increments with the 'depth'
def parse(data):
    score = 0
    open_groups = 0
    parsing_garbage = False
    for character in data:
        # guards, guards!
        if parsing_garbage and not character == ">":
            continue

        if character == "<":
            parsing_garbage = True
        elif character == ">":
            parsing_garbage = False
        elif character == "{":
            open_groups += 1
        elif character == "}":
            score += open_groups
            open_groups -= 1
    return score
   

# Each time you close, {, the score increments with the 'depth'
def parse2(data):
    removed_chars = 0
    parsing_garbage = False
    for character in data:
        # guards, guards!
        if parsing_garbage and not character == ">":
            removed_chars += 1
            continue

        if character == "<":
            parsing_garbage = True
        elif character == ">":
            parsing_garbage = False
    return removed_chars 


def sanitize(data):
    """ Apply escape character to text """
    out = ""
    i = 0
    while i < (len(data)):
        char = data[i]
        if char == "!":
            i += 2
            continue
        out += char
        i += 1
    return out


def getInput():
    f = open("input9.txt","r")
    return f.read()




def solve():
    data = getInput()
    score = parse(sanitize(data))
    print(score)
    print(parse2(sanitize(data)))

solve()
