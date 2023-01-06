import math
import sys
import copy

class Rock:
    def __init__(self, name, height, width, indices):
        self.name = name
        self.height = height
        self.width = width
        self.indices = indices

def trim(row):
    if '#' in row or '-' in row:
        return True
    else:
        return False

def print_cave(cave):
    cave = cave[::-1]
    if len(sys.argv) > 2 and sys.argv[2] == "-v":
        print("\nCave readout:")
        for i, row in enumerate(cave):
            print(len(cave) - 1 - i, "\t", " ".join(row))
        print("\n")

def main(rock_count):
    jets = open("jets.txt", "r").read()

    # Generate initial cave
    width = 7
    empty_cave_row = ['.' for i in range(width)]
    cave_floor = ['-' for i in range(width)]
    cave = [cave_floor.copy()]

    rocks = [
        Rock("horizontal", 1, 4, [0,1,2,3]),
        Rock("cross", 3, 3, [1,7,8,9,15]),
        Rock("l", 3, 3, [2,9,14,15,16]),
        Rock("vertical", 4, 1, [0,7,14,21]),
        Rock("square", 2, 2, [0,1,7,8])
    ]

    placed_rocks = 0
    jet_index = 0

    while placed_rocks < int(rock_count):
        rock = rocks[placed_rocks % len(rocks)]

        stopped = False
        row = 0
        jet_offset = 0
        x_anchor = 2

        for i in range(rock.height + 3):
            cave.append(empty_cave_row.copy())

        height = len(cave) - 1

        rock_coords = []

        fake_cave = copy.deepcopy(cave)

        while not stopped:
            row += 1
            jet_offset = 1 if jets[jet_index % len(jets)] == ">" else -1
            movable = True

            jet_index += 1

            rock_coords = []
            for index in rock.indices:
                y_index = math.floor(index/width)
                y_offset = row + y_index
                horizontal = x_anchor + (index - y_index * width)

                rock_coords.append([height - (y_offset - 1), horizontal])
                target = cave[height - y_offset][horizontal]

                if target == "#" or target == "-":
                    stopped = True

            if not x_anchor + rock.width + jet_offset > width or x_anchor + jet_offset < 0:
                for coords in rock_coords:
                    target = cave[coords[0]][coords[1] + jet_offset]
                    if target == "#":
                        movable = False
                
                if movable:
                    x_anchor += jet_offset
                    for i, coords in enumerate(rock_coords):
                        rock_coords[i][1] = rock_coords[i][1] + jet_offset

            for coords in rock_coords:
                fake_cave[coords[0]][coords[1]] = "#"
            # print_cave(fake_cave)
            fake_cave = copy.deepcopy(cave)

        for coords in rock_coords:
            cave[coords[0]][coords[1]] = "#"

        placed_rocks += 1
    
        cave = list(filter(trim, cave))

    print_cave(cave) 
    print('Rows tall after {} rocks: {}'.format(rock_count, len(cave)))

main(sys.argv[1])
