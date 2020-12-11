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
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
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

func solve2() int {
	grid := getInput()

	c := 0
	for {
		newgrid, changed := timeStep2(grid)
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
		if c == 3 {
			fmt.Printf("grid: %v\n", grid)
			break
		}
		c++
	}

	return 0
}

func timeStep2(grid [][]string) ([][]string, int) {
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
			seats := seatsInLOS(row, col, grid)
			if current == EMPTY && seats == 0 {
				changed++
				cp[row][col] = FULL
			} else if current == FULL && seats >= 5 {
				changed++
				cp[row][col] = EMPTY
			}
		}
	}
	return cp, changed
}

// count seats in Line Of Sight
func seatsInLOS(row, col int, grid [][]string) int {
	out := 0
	// up
	for i := row + 1; i < len(grid)-1; i++ {
		if grid[i][col] == FLOOR {
			continue
		}
		if grid[i][col] == FULL {
			out++
		}
		break
	}
	// down
	for i := row - 1; i >= 0; i-- {
		if grid[i][col] == FLOOR {
			continue
		}
		if grid[i][col] == FULL {
			out++
		}
		break
	}

	// left
	for j := col + 1; j < len(grid[row])-1; j++ {
		if grid[row][j] == FLOOR {
			continue
		}
		if grid[row][j] == FULL {
			out++
		}
		break
	}
	// right
	for j := col - 1; j >= 0; j-- {
		if grid[row][j] == FLOOR {
			continue
		}
		if grid[row][j] == FULL {
			out++
		}
		break
	}

	// diagonals

	j := col + 1
	for i := row + 1; i < len(grid)-1; i++ {
		if j > len(grid[row])-1 {
			break
		}
		if grid[i][j] == FLOOR {
			j++
			continue
		}
		if grid[i][j] == FULL {
			out++
		}
		break
	}

	j = col - 1
	for i := row + 1; i < len(grid)-1; i++ {
		if j < 0 {
			break
		}
		if grid[i][j] == FLOOR {
			j--
			continue
		}
		if grid[i][j] == FULL {
			out++
		}
		break
	}

	j = col + 1
	for i := row - 1; i >= 0; i-- {
		if j > len(grid[row])-1 {
			break
		}
		if grid[i][j] == FLOOR {
			j++
			continue
		}
		if grid[i][j] == FULL {
			out++
		}
		break
	}

	j = col - 1
	for i := row - 1; i >= 0; i-- {
		if j < 0 {
			break
		}
		if grid[i][j] == FLOOR {
			j--
			continue
		}
		if grid[i][j] == FULL {
			out++
		}
		break
	}

	return out
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
	in, _ := ioutil.ReadFile("test.txt")

	out := [][]string{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		out = append(out, strings.Split(line, ""))
	}

	return out
}
