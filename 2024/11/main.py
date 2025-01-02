from typing import List

test_seq = [125, 17]
real_seq = [0, 7, 198844, 5687836, 58, 2478, 25475, 894]




def mod(data: List[int]):
    new_list: List[int] = []
    for num in data:
        if num == 0:
            new_list.append(1)
        elif len(str(num)) % 2 == 0:
            # split in two.. 
            part_one = str(num)[0:len(str(num))//2]
            part_two = str(num)[len(str(num))//2:len(str(num))]
            new_list.append(int(part_one))
            new_list.append(int(part_two))
        else:
            new_list.append(num * 2024)
    return new_list





def solve1(data: List[int]) -> int: 
    for i in range(0, 25):
        data = mod(data)
    return len(data)


if __name__ == '__main__':
    data = real_seq 
    print(solve1(data))
