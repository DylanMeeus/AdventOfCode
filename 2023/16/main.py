from enum import Enum


TEST = False

HEIGHT = 10 if TEST else 100
WIDTH = 10 if TEST else 100


class Direction(Enum):
    UP = 1,
    DOWN = 2,
    LEFT = 3,
    RIGHT = 4


def next_pos(location, direction):
    if direction == Direction.UP:
        return (location[0] - 1, location[1])
    if direction == Direction.DOWN:
        return (location[0] + 1, location[1])
    if direction == Direction.LEFT:
        return (location[0], location[1] - 1)
    if direction == Direction.RIGHT:
        return (location[0], location[1] + 1)


class Beam:
    def __init__(self, location, direction, path = None):
        self.loc = location
        self.dir = direction
        self.active = True 
        H = (self.loc, self.dir)
        if path is not None and H in path:
            self.active = False
        self.path = set() if path == None else path

    def path_changed(self):
        H = (self.loc, self.dir)
        if H in self.path:
            self.active = False
        else:
            self.path.add(H)



    def move(self, G):
        self.path_changed()
        next_tile = next_pos(self.loc, self.dir)

        if self.active == False:
            return

        if next_tile not in G:
            self.active = False

        elif G[next_tile] == ".":
            self.loc = next_tile
        
        elif G[next_tile] == '/':
            self.loc = next_tile

            if self.dir == Direction.UP:
                self.dir = Direction.RIGHT
            elif self.dir == Direction.DOWN:
                self.dir = Direction.LEFT

            elif self.dir == Direction.RIGHT:
                self.dir = Direction.UP
            elif self.dir == Direction.LEFT:
                self.dir = Direction.DOWN


        elif G[next_tile] == '\\':
            self.loc = next_tile

            if self.dir == Direction.UP:
                self.dir = Direction.LEFT
            elif self.dir == Direction.DOWN:
                self.dir = Direction.RIGHT

            elif self.dir == Direction.RIGHT:
                self.dir = Direction.DOWN
            elif self.dir == Direction.LEFT:
                self.dir = Direction.UP


        elif G[next_tile] == '|':
            if self.dir in [Direction.LEFT, Direction.RIGHT]:
                self.active = False
                return {Beam(next_tile, Direction.UP, self.path), Beam(next_tile, Direction.DOWN, self.path)}
            else:
                self.loc = next_tile

        elif G[next_tile] == '-':
            if self.dir in [Direction.UP, Direction.DOWN]:
                self.active = False
                return {Beam(next_tile, Direction.LEFT, self.path), Beam(next_tile, Direction.RIGHT, self.path)}
            else:
                self.loc = next_tile
        else:
            print(f'encountered {G[next_tile]} on {next_tile}')
            return 0
        return None


        # if we can no longer move, we return false

def solve1(G):

    beams = {Beam((0,0), Direction.DOWN)}

    E = set()

    for i in range(0, 1000):
        new_beams = set()
        for b in beams:
            E.add(b.loc)
            if b.active:
                new_beams.add(b)
                output = b.move(G)
                if output is not None:
                    for new_beam in output:
                        new_beams.add(new_beam)

        beams = new_beams
        if len(new_beams) == 0:
            return len(E)

    print('done running')
    print_E(E)
    return len(E)


def generate_start_configs():
    left_top = {Beam((0,0), Direction.DOWN), Beam((0,0), Direction.RIGHT)}
    right_top = {Beam((0,WIDTH), Direction.DOWN), Beam((0,WIDTH), Direction.LEFT)}
    bottom_left = {Beam((HEIGHT,0), Direction.UP), Beam((HEIGHT,0), Direction.RIGHT)}
    bottom_right = {Beam((HEIGHT,WIDTH), Direction.UP), Beam((HEIGHT, WIDTH), Direction.LEFT)}

    configs = [left_top, right_top, bottom_left, bottom_right]

    for i in range(1, WIDTH - 1):
        top = {Beam((0,i), Direction.DOWN)}
        bottom = {Beam((HEIGHT,i), Direction.UP)}
        configs.append(top)
        configs.append(bottom)
    return configs



def solve2(G):


    max_E = 0

    confs = generate_start_configs()


    for config in confs:
        for start_config in config:
            E = set()
            beams = {start_config}
            for i in range(0, 1000):
                new_beams = set()
                for b in beams:
                    E.add(b.loc)
                    if b.active:
                        new_beams.add(b)
                        output = b.move(G)
                        if output is not None:
                            for new_beam in output:
                                new_beams.add(new_beam)

                beams = new_beams
                print(len(E))
                if len(new_beams) == 0:
                    if len(E) > max_E:
                        max_E = len(E)
                    break

    return max_E

def print_E(E):
    out = ""
    for row in range(0, 10):
        for col in range(0, 10):
            H = (row, col)
            if H in E:
                out += "#"
            else:
                out += "."
        out += "\n"
    print(out)


def parse(lines):
    G = {}

    for row, line in enumerate(lines):
        for col, char in enumerate(line):
            loc = (row, col)
            G[loc] = char

    return G

if __name__ == '__main__':
    lines = open('input.txt').read().split('\n')
    G = parse(lines)
    print(f'output : {solve1(G)} ')
    print(f'output2: {solve2(G)} ')
