""" Solution to Advent Of Code 5 """
from functools import reduce
from string import ascii_lowercase
import re

def react(polymer):
    res = reduce(lambda x, y:  x[:-1] if x[-1:] != y and x.lower()[-1:] == y.lower() else x + y, polymer)
    if len(res) == len(polymer):
        return res
    return react(res)

def find_shortest(polymer):
    shortest = len(polymer)
    for c in ascii_lowercase:
        trimmed_polymer = re.sub("[{}{}]".format(c, c.upper()), "", polymer)
        trimmed_len = len(react(trimmed_polymer))
        if trimmed_len < shortest:
            shortest = trimmed_len
    return shortest


if __name__ == '__main__':
    polymer = open('input.txt','r').read().replace("\n","")
    print(len(react(polymer)))
    print(find_shortest(polymer))
