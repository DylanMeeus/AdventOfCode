package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}

type point struct {
	row int
	col int
}

func solve1() int {
	data := getData()
	fmt.Println(data)

	visibleTrees := map[point]bool{}

	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			if visible(row, col, data) {
				visibleTrees[point{row, col}] = true
			}
		}
	}

	fmt.Println(visibleTrees)

	return len(visibleTrees)
}

func solve2() int {

	data := getData()
	fmt.Println(data)

	visibleTrees := map[point]int{}

	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			if visible(row, col, data) {
				visibleTrees[point{row, col}] = countVisible(row, col, data)
			}
		}
	}

	fmt.Println(visibleTrees)

	score := 0
	for _, value := range visibleTrees {
		if value > score {
			score = value
		}
	}

	return score
}

func countVisible(row, col int, data [][]int) int {
	h := data[row][col]

	checkUp := func(r int) (out int) {
		for i := r; i >= 0; i-- {
			if data[i][col] < h {
				out++
			} else if data[i][col] == h {
				return out + 1
			} else {
				return out + 1
			}
		}
		return out
	}

	checkDown := func(r int) (out int) {
		for i := r; i < len(data); i++ {
			if data[i][col] < h {
				out++
			} else if data[i][col] == h {
				return out + 1
			} else {
				return out + 1
			}
		}
		return out
	}

	checkLeft := func(c int) (out int) {
		for i := c; i >= 0; i-- {
			if data[row][i] < h {
				out++

			} else if data[row][i] == h {
				return out + 1
			} else {
				return out + 1
			}
		}
		return out
	}

	checkRight := func(c int) (out int) {
		for i := c; i < len(data[row]); i++ {
			if data[row][i] < h {
				out++
			} else if data[row][i] == h {
				return out + 1
			} else {
				return out + 1
			}
		}
		return out
	}

	return checkUp(row-1) * checkDown(row+1) * checkLeft(col-1) * checkRight(col+1)
}

func visible(row, col int, data [][]int) bool {
	h := data[row][col]

	checkUp := func(r int) bool {
		for i := r; i >= 0; i-- {
			if data[i][col] >= h {
				return false
			}
		}
		return true
	}

	checkDown := func(r int) bool {
		for i := r; i < len(data); i++ {
			if data[i][col] >= h {
				return false
			}
		}
		return true
	}

	checkLeft := func(c int) bool {
		for i := c; i >= 0; i-- {
			if data[row][i] >= h {
				return false
			}
		}
		return true
	}

	checkRight := func(c int) bool {
		for i := c; i < len(data[row]); i++ {
			if data[row][i] >= h {
				return false
			}
		}
		return true
	}

	return checkUp(row-1) || checkDown(row+1) || checkLeft(col-1) || checkRight(col+1)
}

func getData() [][]int {

	f, err := ioutil.ReadFile("./input.txt")
	handleError(err)

	lines := strings.Split(string(f), "\n")

	out := [][]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		cols := strings.Split(line, "")
		inner := make([]int, len(cols))
		for i, char := range cols {
			h, err := strconv.Atoi(char)
			handleError(err)
			inner[i] = h
		}
		out = append(out, inner)
	}

	return out
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
