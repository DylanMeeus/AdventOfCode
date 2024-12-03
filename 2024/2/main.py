

def get_input() -> [int]:
    lines = open('input.txt').read().split("\n")

    out = []

    for line in lines:
        if line == '':
            continue
        out.append(list(map(lambda k: int(k), line.split(" "))))
    return out



def solve2(inp: [int]) -> int:
    count = 0

    for record in inp:
        l = list(map(lambda k: (k[0] - k[1]), zip(record, record[1:])))

        if is_safe(l):
            count += 1
            continue

        if is_safe(l[1:]):
            count += 1
            continue

        asc = record[0] < record[len(record)-1]
        # it should be ascending, so we can iterate through it in this way

        last_checked = record[0]
        error_count = 0 
        for index in range(1,len(record)):
            # check if it conforms.. 
            to_check = record[index]
            if asc:
                # if it starts descending, or if it is out of bounds.. 
                if to_check <= last_checked or to_check > last_checked + 3:
                    #print(f'compared {last_checked} against {to_check}')
                    error_count += 1
                    continue
            if not asc:
                if to_check >= last_checked or to_check < last_checked - 3: 
                    #print(f'compared {last_checked} against {to_check}')
                    error_count += 1
                    continue
            # no errors found with this number
            last_checked = to_check


        if error_count == 1:
            count += 1
        else:
            print(f' unfixable {record} with {error_count} mistakes')



    return count

def is_safe(l) -> bool:
        return len(list(filter(lambda k: k > 0 and k <= 3, l))) == len(l) or len(list(filter(lambda k: k < 0 and k >= -3, l))) == len(l)

    
def solve1(inp: [int]) -> int: 
    count = 0

    for record in inp:
        l = list(map(lambda k: (k[0] - k[1]), zip(record, record[1:])))
        if len(list(filter(lambda k: k > 0 and k <= 3, l))) == len(l) or len(list(filter(lambda k: k < 0 and k >= -3, l))) == len(l):
            count += 1

    return count

if __name__ == '__main__':
    inp = get_input()

    print(solve2([[1,3,2,4,5]]))
    print(solve2([[8,6,4,4,1]]))

    print(solve1(inp))
    print(solve2(inp))

    
