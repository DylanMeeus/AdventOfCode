""" Solution to Advent Of Code 5 """
from functools import reduce

def test(f, polymer):
    return f(polymer) == "dabCBAcaDA"


def react(polymer):
    res = reduce(lambda x, y:  x[:-1] if x[-1:] != y and x.lower()[-1:] == y.lower() else x + y, polymer)
    if len(res) == len(polymer):
        return res
    return react(res)


if __name__ == '__main__':
    polymer = open('input.txt','r').read().replace("\n","")
    print(len(react(polymer)))
