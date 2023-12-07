



def simulate_lap_time(hold_time, distance) -> int:
    """ returns the lap time based on the time the button is held down """ 
    if hold_time == distance:
        return distance + 1 # magic to disqualify

    speed = hold_time
    return distance // speed


def simulate_distance(hold_time,  max_time, distance) -> int:
    remaining_time = max_time - hold_time
    speed = hold_time
    return speed * remaining_time



def solve1(pairs) -> int :
    _mult = 1
    for pair in pairs:
        ways = 0
        for x in range(1, pair[0]):
            lap_distance = simulate_distance(x, pair[0], pair[1])
            is_win = lap_distance > pair[1]
            if is_win:
                ways += 1
        _mult *= ways


    return _mult


if __name__ == '__main__':
    inputs = open('input.txt').read().split("\n")
    times = list(map(lambda k: int(k), filter(lambda k: k != '', inputs[0][len("Time:"):].split(" "))))
    distance = list(map(lambda k: int(k), filter(lambda k: k != '', inputs[1][len("Distance:"):].split(" "))))
    pairs = list(zip(times,distance))
    print(solve1(pairs))
