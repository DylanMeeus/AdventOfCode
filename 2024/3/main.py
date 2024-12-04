import re

def solve1(inp):
    result = 0
    pattern = r'mul\(\d*\,\d*\)'
    matches = re.findall(pattern, inp)
    for match in matches:
        result += eval_mul(match)

        
    return result


def eval_mul(match: str) -> int:
    i = 1
    nums = r'\d*'
    matching_nums = re.findall(nums, match)
    for num in matching_nums:
        if num != '': 
            i *= int(num)
    return i

def solve2(inp):
    result = 0
    pattern = r'do\(\)|don\'t\(\)|mul\(\d*\,\d*\)'
    matches = re.findall(pattern, inp)
    enabled = True

    for match in matches:
        if match == "don't()":
            enabled = False
        elif match == "do()":
            enabled = True
        else:
            if enabled:
                result += eval_mul(match)


    return result



def get_input() -> str: 
    return open('input.txt').read()


if __name__ == '__main__':
    inp = get_input()
    print(solve1(inp))
    print(solve2(inp))
