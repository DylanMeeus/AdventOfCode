

class bot:
    def __init__(self, pos, rad):
        self.pos = pos
        self.rad = rad



def solve(bots):
    # best radius bot
    bot = max(bots, key=lambda b: b.rad)
    in_range = 0
    for b in bots:
        dist = distance(bot,b)
        if dist <= bot.rad:
            in_range += 1
    return in_range


def distance(bota, botb):
    posa = bota.pos
    posb = botb.pos
    return abs(posa[0] - posb[0]) + abs(posa[1] - posb[1]) + abs(posa[2] - posb[2])


def parse(lines):
    bots = []
    for line in lines:
        if line == "":
            continue
        parts = line.split(" ")
        position = parts[0][4:]
        sane = position.replace("<","").replace(">","")[:-1]
        sane = sane.split(",")
        pos = (int(sane[0]), int(sane[1]), int(sane[2]))
        radius = int(parts[1][2:])
        bots.append(bot(pos,radius))
    return bots

if __name__ == '__main__':
    lines = open("input.txt",'r').read().split("\n") 
    print(solve(parse(lines)))


