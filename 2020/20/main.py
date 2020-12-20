import math
import copy


class point:
    def __init__(self, x,y):
        self.x = x
        self.y = y

def rotate90(tile):
    rotated = zip(*tile[::-1])
    l = list(rotated)
    out = []
    for x in l:
        out.append(''.join(x))
    return out





def solve1():
    img = get_input()

    test = img['2311']

    # create a rectangular image
    size = int(math.sqrt(len(img)))
    
    for tile in img:
        backtrack(tile, img, {}, 0, size)




# backtrack to find aligned tiles?
def backtrack(current, tiles, grid, current_idx, size):

    # we map the current idx to the size and then look for neighbours?

    if len(tiles) == 0:
        return True

    # we have to find the right, top, bottom and left border?
    img = tiles[current]
    if current_idx:
        grid[point(0,0)] = img
    else:
        # find position in grid
        if current_idx == 1:
            # let's just simplify it for this



    copy_tiles = copy.deepcopy(tiles)
    del copy_tiles[current]

    for remainder_tile in copy_tiles:
        backtrack(remainder_tile, copy_tiles, grid, current_idx + 1)








def get_input():
    file = open("input_test.txt", "r")
    input = (file.read())
    m = {}
    for tile in (input.split("\n\n")): 
        tile_parts = tile.split("\n")
        title = tile_parts[0].split(" ")[1][:-1]
        tiles = tile_parts[1:]
        m[title] = tiles

    return m




if __name__ == '__main__':
    solve1()
