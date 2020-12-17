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

// layer is a 2-d slice of the world?
type layer struct {
	idx  int // layer index? Maybe store the layers in a LinkedList??
	grid [][]bool
}

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	baseLayer := getInput()

	baseLayer = pad(baseLayer)
	fmt.Printf("%v\n", baseLayer)
	return 1

	// create the first layer?
	layers := make([]layer, 1+ITERATIONS)
	layers[len(layers)/2] = baseLayer

	for i := 0; i < ITERATIONS; i++ {
		layers = timeStep(layers)
		fmt.Printf("%v\n", layers)
	}

	return 1
}

// timeStep does one 'tick' to the next state
func timeStep(layers []layer) []layer {
	// work on a copy?
	cl := copyLayers(layers)

	for _, layer := range cl {
		for row := range layer.grid {
			for col := range layer.grid[row] {
				count := countNeighbours(row, col, layer.idx, layers)
				fmt.Printf("%v\n", count)
			}
		}
	}
	return cl
}

// add some padding to the grid
func pad(input layer) layer {
	padding := make([][]bool, 3)
	rowPadding := make([]bool, 3)
	for row := range input.grid {
		input.grid[row] = append(rowPadding, input.grid[row]...)
		input.grid[row] = append(input.grid[row], rowPadding...)
	}
	// should we extend them appropriately?
	input.grid = append(padding, input.grid...)
	input.grid = append(input.grid, padding...)
	return input
}

func countNeighbours(x, y, z int, layers []layer) int {

	return 0
}

// do a deep-copy
func copyLayers(layers []layer) []layer {
	out := []layer{}

	for _, l := range layers {
		newLayer := layer{}
		newLayer.grid = make([][]bool, len(l.grid))
		for i := 0; i < len(l.grid); i++ {
			newLayer.grid[i] = make([]bool, len(l.grid[i]))
			for j := 0; j < len(l.grid[i]); j++ {
				newLayer.grid[i][j] = l.grid[i][j]
			}
		}
		out = append(out, newLayer)

	}

	return out
}

// getInput returns the input as a layer
func getInput() layer {
	in, _ := ioutil.ReadFile("test.txt")
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
