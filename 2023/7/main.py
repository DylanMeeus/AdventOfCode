from enum import Enum 

# type x = (card, value)


card_order = {
        'A': 0, 
        'K':1, 
        'Q':2, 
        'J':3, 
        'T':4 , 
        '9':5,
        '8':6,
        '7':7,
        '6':8,
        '5':9,
        '4':10,
        '3':11,
        '2':12
        }


class hand_type(Enum):
    FIVE_OF_KIND = 6
    FOUR_OF_KIND = 5
    FULL_HOUSE = 4
    THREE_OF_KIND = 3
    TWO_PAIR = 2
    ONE_PAIR = 1
    HIGH_CARD = 0

    def __lt__(self, other):
        if self.__class__ is other.__class__:
            return self.value < other.value





class entry:
    def __init__(self, cards, bid):
        self.cards = cards
        self.bid = bid
        self.type = self.derive_type()

    def derive_type(self):
        card_map = {}
        for card in self.cards:
            if card not in card_map:
                card_map[card] = 1
            else:
                card_map[card] += 1
        card_values = list(card_map.values())
        if len(card_map) == 1:
            return hand_type.FIVE_OF_KIND
        if len(card_map) == 2 and (card_values[0] ==  1 or card_values[1] == 1):
            return hand_type.FOUR_OF_KIND
        if len(card_map) == 2 and (card_values[0] ==  3 or card_values[1] == 3):
            return hand_type.FULL_HOUSE
        if len(card_map) == 3:
            if 3 in card_values:
                return hand_type.THREE_OF_KIND
            if 2 in card_values:
                return hand_type.TWO_PAIR
        if len(card_map) == 4:
            if 2 in card_values:
                return hand_type.ONE_PAIR
        return hand_type.HIGH_CARD


    def cmp_cards(self, other):
        for idx, card in enumerate(self.cards):
            if card == other.cards[idx]:
                continue
            return card_order[card] > card_order[other.cards[idx]]
        return False

    def __lt__(self, other):
        if self.type == other.type:
            return self.cmp_cards(other)
        return self.type < other.type

    def __str__(self):
        return f'{self.cards}'

    def __repr__(self):
        return f'{self.cards}' 


def solve1(entries):
    sorted_entries = sorted(entries)
    _sum = 0
    for idx, entry in enumerate(sorted_entries):
        _sum += (entry.bid * (idx+1))
        print(f'{entry.type} - {entry.cards}')
    return _sum
    

def parse(lines):
    entries = []
    for line in lines:
        if line == '':
            continue
        parts = line.split(' ')
        entries.append(entry(parts[0], int(parts[1])))
    return entries



if __name__ == '__main__':
    lines = open('input.txt').read().split('\n')
    entries = parse(lines)
    print(solve1(entries))

    e1 = entry('AAA66', 0)
    print(e1.type)





