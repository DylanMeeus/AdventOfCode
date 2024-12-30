from typing import List, Tuple

def get_input() -> str:
    return open('input.txt').read()


def compress(l: List[str]) -> List[str]:
    """ compress using a two-pointer approach """ 
    start = 0
    end = len(l) - 1

    while start < end: 
        if l[start] != '.':
            start += 1
            continue
        else:
            if l[end] == '.':
                end -= 1
                continue
            else:
                # lwap them.. 
                copy = l[start]
                l[start] = l[end]
                l[end] = copy
                start += 1

    return l

def expand(s: str) -> List[str]:
    out: List[str] = []
    idx = 0
    for i, freq in enumerate(s):
        if not freq.isdigit():
            continue
        if i % 2 == 0:
            for i in range(0, int(freq)):
                out.append(str(idx))
            idx += 1
        else:
            for i in range(0, int(freq)):
                out.append(".")
    return out
        

def checksum(data: List[str]) -> int:
    out = 0
    for idx, file_id in enumerate(data):
        if file_id == '.':
            None
        else:
            out += (idx * int(file_id))
    return out

def solve1(data) -> int:
    expanded = expand(data)
    compressed = compress(expanded)
    return checksum(compressed)


def find_free_space(space_needed: int, max_loc: int, data: List[str]) -> Tuple[int, int]:
    """ return the index of where to insert the data """ 
    for i in range(0, max_loc):
        if data[i] == '.':
            if space_needed == 1:
                return i, i
            space_counter = 1
            for j in range(i+1, max_loc):
                if j < len(data) and data[j] == '.':
                    space_counter += 1
                    if space_counter == space_needed:
                        return i, j 
                else:
                    break
    return -1, -1






def whole_file_compression(data: List[str]) -> List[str]:
    # count how many instances there are for each ID 
    freq_map = {}
    for char in data:
        if char not in freq_map:
            freq_map[char] = 0
        freq_map[char] += 1

    files_to_place = list(reversed(freq_map.keys()))

    for i, file_id in enumerate(files_to_place):
        print(f'compressed {i} of {len(files_to_place)} files...')
        space_needed = freq_map[file_id]
        #print(f'checking space needed for {file_id} - {space_needed}')
        start_idx, end_idx = find_free_space(space_needed, data.index(file_id), data)
        if start_idx == -1:
            continue
        # now mutate the data list to place the data elements.. 
        data = list(map(lambda k: '.' if k == file_id else k, data))
        for i in range(start_idx, end_idx+1):
            data[i] = file_id

    return data


def solve2(data) -> int:
    expanded = expand(data)
    compressed = whole_file_compression(expanded)
    return checksum(compressed)


if __name__ == '__main__':
    data = get_input()
    print(solve1(data))
    print(solve2(data))
