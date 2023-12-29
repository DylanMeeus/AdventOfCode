from enum import Enum


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
        self.path = set() if path is None else path

    def path_changed(self):
        H = (self.loc, self.dir)
        if H in self.path:
            self.active = False
        else:
            self.path.add(H)



    def move(self, G):
        self.path.add(self.loc)
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
        print(len(E))
        if len(new_beams) == 0:
            return len(E)

    return len(E)


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
