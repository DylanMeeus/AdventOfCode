package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	ACTIVE     = true
	INACTIVE   = false
	ITERATIONS = 6
)

/*
var directions = []point{
	{0, 0, 1},
	{0, 1, 0},
	{1, 0, 0},
	{1, 0, 1},
	{0, 1, 1},
	// and now negative
	{0, 0, -1},
	{0, -1, 0},
	{-1, 0, 0},
	{-1, 0, -1},
	{0, -1, -1},
	//
	{1, 0, -1},
	{0, 1, -1},
	{-1, 0, 1},
	{0, -1, 1},
}
*/

type point struct {
	x, y, z int
}

type layer struct {
	grid [][]bool
}

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	baseLayer := getInput()

	field := map[point]bool{}

	n := len(baseLayer.grid) + (ITERATIONS)

	for z := -n; z < n; z++ {
		for y := -n; y < n; y++ {
			for x := -n; x < n; x++ {
				field[point{x, y, z}] = false
			}
		}
	}

	for row := range baseLayer.grid {
		for col := range baseLayer.grid[row] {
			field[point{row, col, 0}] = baseLayer.grid[row][col]
		}
	}

	// we need space to grow into :-)

	for i := 0; i < ITERATIONS; i++ {
		field = timeStep(field)
	}

	activos := 0

	fmt.Printf("%v\n", field)
	for _, v := range field {
		if v {
			activos++
		}
	}

	return activos
}

// timeStep does one 'tick' to the next state
func timeStep(field map[point]bool) map[point]bool {
	// work on a copy?
	cf := copyField(field)

	for point, state := range field {
		n := countNeighbours(point, field)
		if state == ACTIVE && (n != 2 && n != 3) {
			cf[point] = INACTIVE
		} else if state == INACTIVE && n == 3 {
			cf[point] = ACTIVE
		}
	}

	return cf
}

func countNeighbours(p point, field map[point]bool) (out int) {
	// any of their coordinates differ by at most 1
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				neighbour := point{x, y, z}
				if neighbour != p && field[neighbour] {
					out++
				}

			}
		}
	}
	return
}

func copyField(input map[point]bool) map[point]bool {
	out := map[point]bool{}
	for k, v := range input {
		out[k] = v
	}
	return out
}

// getInput returns the input as a layer
func getInput() layer {
	in, _ := ioutil.ReadFile("input.txt")
	l := layer{}
	lines := strings.Split(string(in), "\n")
	l.grid = make([][]bool, (len(lines) - 1))
	for i, line := range lines {
		if line == "" {
			continue
		}
		states := strings.Split(line, "")
		l.grid[i] = make([]bool, len(states))
		for j, state := range states {
			if string(state) == "#" {
				l.grid[i][j] = ACTIVE
			} // else inactive is default
		}
	}
	return l
}
