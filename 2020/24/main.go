package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const (
	WHITE = false
	BLACK = true
)

var (
	combined = map[string]bool{
		"se": true,
		"sw": true,
		"ne": true,
		"nw": true,
	}
)

// let's encode a hexagonal grid as a grid of cubes
type point struct {
	x, y, z int
}

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve2() int {
	input := getInput()
	floor := map[point]bool{}

	for _, line := range input {
		x, y, z := findPosition(line)
		floor[point{x, y, z}] = !floor[point{x, y, z}]
		//floor[point{x, y, z}] = BLACK
	}

	// is this the starting configuration?

	out := 0
	for day := 0; day < 100; day++ {
		fmt.Printf("day: %v\n", day)
		floor = timeStep(floor)
	}

	for _, v := range floor {
		if v == BLACK {
			out++
		}
	}
	return out
}

func abs(i int) int {
	return int(math.Abs(float64(i)))

}

func timeStep(input map[point]bool) map[point]bool {

	// always increase the size smartly by one?
	SIZE := 100
	out := map[point]bool{}
	for x := -SIZE; x < SIZE; x++ {
		for y := -SIZE; y < SIZE; y++ {
			for z := -SIZE; z < SIZE; z++ {
				p := point{x, y, z}
				c := input[p]
				n := countNeighbours(input, p)
				out[p] = input[p]
				if (c == BLACK && n == 0) || (c == BLACK && n > 2) {
					out[p] = WHITE
				}

				if c == WHITE && n == 2 {
					out[p] = BLACK
				}
			}
		}
	}
	return out

}

// count adjacent black tiles
func countNeighbours(floor map[point]bool, position point) int {
	// x y z points
	var directions = []point{
		{0, 1, -1}, // NW
		{1, 0, -1}, // NE
		{1, -1, 0}, // E
		{0, -1, 1}, // SW
		{-1, 0, 1}, // SW
		{-1, 1, 0}, // W
	}

	out := 0
	for _, dir := range directions {
		pos := point{
			x: position.x + dir.x,
			y: position.y + dir.y,
			z: position.z + dir.z,
		}
		if floor[pos] == BLACK {
			out++
		}
	}
	return out

}

func solve1() int {

	input := getInput()

	floor := map[point]bool{}

	for _, line := range input {
		x, y, z := findPosition(line)
		floor[point{x, y, z}] = !floor[point{x, y, z}]
	}

	// and now we need to find the black side up..
	out := 0

	for _, v := range floor {
		if v == BLACK {
			out++
		}
	}

	return out
}

func findPosition(tiles string) (int, int, int) {
	x := 0
	y := 0
	z := 0

	var traverse func(input string)
	traverse = func(input string) {
		if len(input) == 0 {
			return
		}

		var head, tail string
		if len(input) > 1 {
			head = input[0:2]
			if len(input) > 2 {
				tail = input[2:]
			}
		} else {
			head = string(input[0])
			tail = input[1:]
		}

		if combined[head] {
			switch head {
			case "nw":
				{
					z--
					y++
				}
			case "ne":
				{
					x++
					z--
				}
			case "sw":
				{
					x--
					z++
				}
			case "se":
				{
					z++
					y--
				}
			}
			traverse(tail)
		} else {
			head = string(input[0])
			tail = input[1:]
			switch head {
			case "e":
				{
					x++
					y--
				}
			case "w":
				{
					x--
					y++
				}
			}
			traverse(tail)
		}
	}

	traverse(tiles)

	return x, y, z

}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")

	filter := func(s []string) (out []string) {
		for _, x := range s {
			if x != "" {
				out = append(out, x)
			}
		}
		return
	}

	return filter(strings.Split(string(in), "\n"))
}
