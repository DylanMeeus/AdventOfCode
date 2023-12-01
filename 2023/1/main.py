



digit_words = ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']

word_to_digit =  { None: -1, 'zero': '0', 'one' : '1', 'two' : '2', 'three': '3', 'four' : '4', 'five' : '5', 'six': '6', 'seven': '7', 'eight': '8', 'nine': '9'}


not_found_pos = 10e6
not_found_neg = -10e6

def first_digit_idx(word):
    for i in range(0, len(word)):
        if word[i].isdigit():
            return (i, word[i])
    return (not_found_pos, '')


def last_digit_idx(word):
    for i in range(0, len(word)):
        last_idx = len(word) - i - 1
        if word[last_idx].isdigit():
            return (last_idx, word[last_idx])
    return (not_found_neg, '')

def first_word_idx(word):
    idx = not_found_pos
    match = None
    for dw in digit_words:
        word_idx = word.find(dw)
        if word.find(dw) != -1 and word_idx < idx:
            idx = word_idx
            match = dw
    return (idx, word_to_digit[match])



def last_word_idx(word):
    idx = not_found_neg
    match = None
    for dw in digit_words:
        word_idx = word.rfind(dw)
        if word.find(dw) != -1 and word_idx > idx :
            idx = word_idx
            match = dw
    return (idx, word_to_digit[match])


def solve2():
    f = open('input.txt')

    lines = f.read().split("\n")

    _sum = 0
    for word in lines:
        first = None
        last = None
        if word == '':
            continue
        fdi, first_digit = first_digit_idx(word)
        fwi, first_word = first_word_idx(word)
        first = first_digit if fdi < fwi else first_word

        ldi, last_digit = last_digit_idx(word)
        lwi, last_word = last_word_idx(word)
        last = last_digit if ldi > lwi else last_word

        print(f'{lwi} vs {ldi}')

        print(f'{word} == {first+last}')
        
        _sum = _sum + int(first + last)

    return _sum



def solve1():
    f = open('input.txt')
    lines = f.read().split("\n")
    print(lines)


    _sum = 0
    for word in lines:
        if len(word) == 0:
            continue
        first = None
        last = None
        for i in range(0, len(word)):
            if word[i].isdigit() and first is None:
                first = word[i]
            last_idx = len(word) - i - 1
            if word[last_idx].isdigit() and last is None:
                last = word[last_idx]
            if first is not None and last is not None:
                print(f'{first} - {last}')
                break
        _sum = _sum + int(first+last)
    return _sum



if __name__ == '__main__':
    print(solve2())



