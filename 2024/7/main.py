from dataclasses import dataclass


@dataclass(frozen=True)
class Eq:
    target: int
    components: [int]


def get_input() -> [Eq]:
    lines = open('input.txt').read().split('\n')
    out = []
    for line in lines:
        if line == '':
            continue
        parts = line.split(':')
        target = int(parts[0])
        components = list(map(lambda k: int(k), parts[1].strip().split(' ')))
        out.append(Eq(target, components))
    return out

def can_solve(t, c, acc) -> bool:
    if len(c) == 0 and acc == t: 
        return True

    if len(c) == 0 or acc > t:
        return False

    to_consume = c[0]
    remainder = c[1:]
    return can_solve(t, remainder, acc + to_consume) or can_solve(t, remainder, acc * to_consume)



def can_solve2(t, c, acc) -> bool:

    if len(c) == 0 and acc == t: 
        return True

    if len(c) == 0 or acc > t:
        return False

    to_consume = c[0]
    remainder = c[1:]

    str_concat = str(acc) + str(to_consume)
    return can_solve2(t, remainder, acc + to_consume) or can_solve2(t, remainder, acc * to_consume) or can_solve2(t, remainder, int(str_concat))



def solve1(data) -> int:
    result = 0
    for eq in data:
        if can_solve(eq.target, eq.components[1:], eq.components[0]):
            result += eq.target
    return result

def solve2(data) -> int:
    result = 0
    for eq in data:
        if can_solve2(eq.target, eq.components[1:], eq.components[0]):
            result += eq.target
    return result

if __name__ == '__main__':
    data = get_input()
    print(solve1(data))
    print(solve2(data))
