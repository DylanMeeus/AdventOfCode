
def solve1(tups):
    _sum = 0
    for tup in tups:
        wins, cards = tup[0], tup[1]
        points = 0
        for w in wins:
            for c in cards:
                if c == w:
                    if points == 0:
                        points = 1
                    else:
                        points *= 2
        _sum += points
    return _sum



def solve2(tups):
    card_play_counter = {}
    for idx, tup in enumerate(tups):
        card_play_counter[idx] = 1

    for idx, tup in enumerate(tups):
        # how often do we have to play this card?
        rounds = card_play_counter[idx]
        wins, cards = tup[0], tup[1]
        for r in range(rounds): 
            matching = 0
            for w in wins:
                for c in cards:
                    if c == w:
                        matching += 1
            for i in range(idx + 1, idx + matching + 1):
                if i in card_play_counter:
                    card_play_counter[i] += 1

    return sum(card_play_counter.values())




if __name__ == '__main__':
    lines = open('input.txt').read().split("\n")
    # parse to ([int],[int])
    tups = []
    for line in lines:
        if line == "":
            continue
        parts = line.split(":")[1]
        parts = parts.split("|")
        wins, cards = parts[0], parts[1]
        wins = list(map(lambda k: int(k), list(filter(lambda k: k != "" ,wins.split(" ")))))
        cards = list(map(lambda k: int(k), list(filter(lambda k: k != "", cards.split(" ")))))
        tups.append((wins,cards))
    print(solve1(tups))
    print(solve2(tups))


