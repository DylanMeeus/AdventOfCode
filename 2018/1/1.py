
def solve(nums):
    return sum(nums)

def cycle(nums):
    i = 0
    while True:
        yield nums[i]
        i = i + 1 if i + 1 < len(nums) else 0

def solve2(nums):
    total = 0
    generator = cycle(nums)
    pnums = {0}
    for i in generator:
        total += i
        if total in pnums:
            return total
        pnums.add(total)




if __name__ == '__main__':
    strnums = open('input.txt', 'r').read().split("\n")
    ints = list(map(lambda x: int(x), filter(lambda y: y != '', strnums)))
    print(solve(ints))
    print(solve2(ints))
    
