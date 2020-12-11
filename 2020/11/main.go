package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	EMPTY = "L"
	FULL  = "#"
	FLOOR = "."
)

func main() {
	fmt.Println("vim-go")
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	grid := getInput()

	for {
		newgrid, changed := timeStep(grid)
		grid = newgrid
		count := 0
		if changed == 0 {
			for _, row := range grid {
				for _, col := range row {
					if col == FULL {
						count++
					}
				}
			}
			return count
		}
	}

	return 0
}

func timeStep(grid [][]string) ([][]string, int) {
	cp := make([][]string, len(grid))
	for i := 0; i < len(grid); i++ {
		cp[i] = make([]string, len(grid[i]))
		for j, val := range grid[i] {
			cp[i][j] = val
		}
	}
	changed := 0
	for row := range grid {
		for col, current := range grid[row] {
			seats := countOccupied(row, col, grid)
			if current == EMPTY && seats == 0 {
				changed++
				cp[row][col] = FULL
			} else if current == FULL && seats >= 4 {
				changed++
				cp[row][col] = EMPTY
			}
		}
	}
	return cp, changed
}

func countOccupied(row, col int, grid [][]string) (out int) {
	if row > 0 && col > 0 {
		if grid[row-1][col-1] == FULL {
			out++
		}
	}
	if row > 0 && col < len(grid[row])-1 {
		if grid[row-1][col+1] == FULL {
			out++
		}
	}
	if row < len(grid)-1 && col > 0 {
		if grid[row+1][col-1] == FULL {
			out++
		}
	}
	if row < len(grid)-1 && col < len(grid[row])-1 {
		if grid[row+1][col+1] == FULL {
			out++
		}
	}

	if row > 0 {
		if grid[row-1][col] == FULL {
			out++
		}
	}
	if row < len(grid)-1 {
		if grid[row+1][col] == FULL {
			out++
		}
	}
	if col > 0 {
		if grid[row][col-1] == FULL {
			out++
		}
	}
	if col < len(grid[row])-1 {
		if grid[row][col+1] == FULL {
			out++
		}
	}

	return
}

func getInput() [][]string {
	in, _ := ioutil.ReadFile("input.txt")

	out := [][]string{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		out = append(out, strings.Split(line, ""))
	}

	return out
}
