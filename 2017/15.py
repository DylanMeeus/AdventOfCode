# Solution to day 15 of AdventOfCode










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




solve()
