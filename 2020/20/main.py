import math
import copy


top = "top"
bottom = "bottom"
right = "right"
left = "left"


def rotate90(tile):
    rotated = zip(*tile[::-1])
    l = list(rotated)
    out = []
    for x in l:
        out.append(''.join(x))
    return out


def flip(tile):
    """ flip the columns .. """
    c = []
    for line in tile:
        c.append(line[::-1])
    return c
    


def solve1():
    img = get_input()

    a = img['2311']
    b = img['1951']
    
    # create a rectangular image
    size = int(math.sqrt(len(img)))
    
    for tile in img:
        backtrack(tile, img, {}, 0, size)



def get_borders(tile):
    borders = {}
    borders[top] = tile[0]
    borders[bottom] = tile[len(tile)-1]
    if borders[bottom] == "":
        borders[bottom] = tile[len(tile)-2]



    left_side = ""
    right_side = ""
    for line in tile:
        if line == "":
            continue
        left_side += line[0]
        right_side += line[-1:]

    borders[left] = left_side
    borders[right] = right_side


    return borders


def are_aligned(my_border, other_border, side):
    # side is relative to my border
    if side == "left":
        if my_border["left"] == other_border["right"]:
            return True
    if side == "right":
        if my_border["right"] == other_border["left"]:
            return True
    if side == "bottom":
        if my_border["bottom"] == other_border["top"]:
            return True
    if side == "top":
        if my_border["top"] == other_border["bottom"]:
            return True

    return False


def fits(position, tile, grid):
    """ true if it fits """
    if len(tile) == 0:
        return

    directions = { "bottom": (1,0), "top": (-1,0), "left": (0,1), "right": (0, -1)}

    ok = False


    my_borders = get_borders(tile)
    for side, dir in directions.items(): 
        new_pos = (position[0] + dir[0], position[1] + dir[1])
        if new_pos in grid:
            other_borders = get_borders(grid[new_pos])
            if are_aligned(my_borders, other_borders, side):
                return True

    return False



# backtrack to find aligned tiles?
def backtrack(current, tiles, grid, current_idx, size):

    # we map the current idx to the size and then look for neighbours?

    if len(tiles) == 0:
        return True

    # we have to find the right, top, bottom and left border?
    img = tiles[current]
    if current_idx == 0:
        grid[(0,0)] = img
        copy_tiles = copy.deepcopy(tiles)
        del copy_tiles[current]

        for remainder_tile in copy_tiles:
            backtrack(remainder_tile, copy_tiles, grid, current_idx + 1, size)
            break
    else:
        # find position in grid
        if current_idx == 1:
            # let's just simplify it for this
            for tileID in tiles:
                tile = tiles[tileID]
                for i in range(0,4):
                    tile = rotate90(tile)
                    if fits((0,1), tile, grid):
                        print("fits here")
                    if fits((0,1), flip(tile), grid):
                        print("fits flipped")




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
