# Solution to day 15 of AdventOfCode




def solve2():
    A = 277
    B = 349
    A_mult = 16807
    B_mult = 48271
    divisor = 2147483647

    mask = 0xffff
    As = []
    Bs = []
    limit = int(5*(10**6))
    while len(As) < limit or len(Bs) < limit:
        if len(As) < limit:
            A = next_val(A,A_mult,divisor)
            if A % 4 == 0:
                As.append(A)
        if len(Bs) < limit: 
            B = next_val(B,B_mult,divisor)
            if B % 8 == 0:
                Bs.append(B)
    print(find_matches(As,Bs,mask))


def find_matches(Xs,Ys,mask):
    matches = 0
    for i in range(len(Xs)):
        if Xs[i] & mask == Ys[i] & mask:
            matches += 1
    return matches



def solve():
    A = 277
    B = 349
    A_mult = 16807
    B_mult = 48271
    divisor = 2147483647
    
    # check 'm
    matches = 0
    mask = 0xffff
    for i in range(int(40*(10**6))):
        A = next_val(A,A_mult,divisor)
        B = next_val(B,B_mult,divisor)
        if A & mask == B & mask:
            matches += 1
    print(matches)




def next_val(init, multiplier, divisor):
    return (init * multiplier) % divisor




#solve()
solve2()
