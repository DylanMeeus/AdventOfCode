
from functools import reduce




def find_next(values):
    derivatives = []
    for i in range(1, len(values)):
        derivatives.append(values[i] - values[i-1])
    if len(list(filter(lambda k: k == 0, derivatives))) == len(derivatives):
        return values[0]
    return values[len(values)-1] + find_next(derivatives)

def solve1(readings):
    _sum = 0
    for reading in readings:
        _sum += find_next(reading)
    return _sum


def solve2(readings):
    _sum = 0 
    for reading in readings:
        reading.reverse()
        _sum += find_next(reading)
    return _sum




if __name__ == '__main__':
    lines = open('input.txt').read().split('\n')[:-1]


    readings = []
    for line in lines:
        readings.append(list(map(lambda k: int(k), line.split(' '))))


    print(solve1(readings))
    print(solve2(readings))

