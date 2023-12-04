import re


class number():
    def __init__(self, value, row, start_col, end_col):
        self.value = value
        self.start = start_col
        self.end = end_col
        self.row = row

    def is_adjacent_to_sign(self, sign) -> bool: 
        for x in range(self.start, self.end):
            if self._has_sign_adjacent(x, sign):
                return True
        return False

    def _has_sign_adjacent(self, col, sign) -> bool:
        # below
        if self.row + 1 == sign.row and col == sign.col:
            return True
        elif self.row + 1 == sign.row and col + 1 == sign.col:
            return True
        elif self.row + 1 == sign.row and col - 1 == sign.col:
            return True
        #above
        elif self.row - 1 == sign.row and col == sign.col:
            return True
        elif self.row - 1 == sign.row and col - 1 == sign.col:
            return True
        elif self.row - 1 == sign.row and col + 1 == sign.col:
            return True
        # col checks
        elif self.row  == sign.row and col + 1 == sign.col:
            return True
        elif self.row  == sign.row and col - 1 == sign.col:
            return True
        else:
            return False



class sign():
    def __init__(self, sign, row, col):
        self.sign = sign
        self.row = row
        self.col = col 


def solve1(lines) -> int:
    # find all numbers
    # write regex to find all numbers - find() returns idx, and we do this for all of them
    # save the number, with idx for each digit
    
    # find all symbols
    # save the idx of the symbol
    numbers, signs = find_numbers_and_signs(lines)
    _sum = 0

    for num in numbers:
        for s in signs:
            if num.is_adjacent_to_sign(s):
                _sum += num.value
                # goto num in numbers
                break


    return _sum


def solve2(lines) -> int:
    numbers, signs = find_numbers_and_signs(lines)
    _sum = 0
    for s in signs:
        if s.sign != "*":
            continue

        collected = []
        for num in numbers:
            if num.is_adjacent_to_sign(s):
                collected.append(num)
        if len(collected) == 2:
            _sum += (collected[0].value * collected[1].value)

    return _sum



def find_numbers_and_signs(lines) -> ([number], [sign]):
    numbers = []
    signs = []
    for row, line in enumerate(lines):
        if line == "":
            continue

        # find numbers
        number_str = re.finditer(r'\d+', line)
        matches = [(m.group(), m.span()) for m in number_str]
        for match in matches:
            num = number(int(match[0]), row, match[1][0], match[1][1])
            numbers.append(num)
        
        for idx, char in enumerate(line):
            if not char.isdigit() and char is not ".":
                signs.append(sign(char, row, idx))

    return (numbers, signs)



if __name__ == '__main__':
    lines  = open('input.txt').read().split("\n")
    print(solve1(lines))
    print(solve2(lines))
    


