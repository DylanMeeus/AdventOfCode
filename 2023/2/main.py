



class throw:
    def __init__(self, r=0, g=0, b=0):
        self.r = r
        self.g = g
        self.b = b


class game:
    def __init__(self, idx):
        self.id = idx
        self.throws = []
        

    def add_throw(self, t):
        self.throws.append(t)


def solve2(games):
    _sum = 0
    for game in games:
        _sum += cube_power(game)
    return _sum

def solve1(games):
    max_red = 12
    max_green = 13
    max_blue = 14

    _sum = 0
    for game in games:
        if is_possible(game, max_red, max_green, max_blue):
            _sum += game.id
    return _sum

        

def is_possible(g, max_r, max_g, max_b):
    count_r, count_g, count_b = 0,0,0
    for throw in g.throws:
        if throw.r > max_r or throw.g > max_g or throw.b > max_b:
            return False

    return True

def cube_power(g):
    min_r, min_g, min_b = 1, 1, 1 
    for throw in g.throws:
        if throw.r > min_r and throw.r != 0:
            min_r = throw.r
        if throw.g > min_g and throw.g != 0:
            min_g = throw.g
        if throw.b > min_b and throw.b != 0:
            min_b = throw.b

    return min_r * min_g * min_b
    


def parse(lines):
    games = []
    for x, line in enumerate(lines):
        if line == "":
            continue
        inputs = line.split(":")[1]
        throws = inputs.split(";")

        _game = game(x+1)
        for t in throws:
            dice = t.split(",")
            r, g, b = 0, 0, 0
            for d in dice:
                parts = d.strip().split(" ")
                num, rgb = int(parts[0]), parts[1]

                if rgb == "red":
                    r = num
                elif rgb == "green":
                    g = num
                elif rgb == "blue":
                    b = num
                else:
                    print(f'fubarred on {rgb}')
                    exit(-1)
            _game.add_throw(throw(r,g,b))
        games.append(_game)
    return games
        

            

if __name__ == '__main__':
    lines = open('input.txt').read().split("\n")
    games = parse(lines)
    print(solve1(games))
    print(solve2(games))
