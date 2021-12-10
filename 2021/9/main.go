package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type heightmap [][]int

func (h heightmap) valueAtPoint(p point) (int, error) {
	if p.row < 0 || p.row >= len(h) {
		return -1, fmt.Errorf("these are not the coordinates you are looking for")
	}

	if p.col < 0 || p.col >= len(h[p.row]) {
		return -1, fmt.Errorf("these are not the coordinates you are looking for")
	}

	return h[p.row][p.col], nil
}

type point struct{ row, col int }

func main() {
	fmt.Printf("%v\n", solve())
}

func getData() heightmap {
	in, _ := ioutil.ReadFile("./input.txt")

	rows := [][]int{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}

		cols := []int{}
		for _, cellValue := range strings.Split(line, "") {
			i, err := strconv.Atoi(cellValue)
			if err != nil {
				panic(err)
			}
			cols = append(cols, i)
		}
		rows = append(rows, cols)
	}

	return heightmap(rows)
}

func solve() int {
	data := getData()
	lowPoints := getLowPoints(data)
	sum := 0
	for _, lp := range lowPoints {
		value, _ := data.valueAtPoint(lp)
		sum += value + 1
	}
	return sum
}

// getLowPoints returns all points for the heightmap
// which are lower than their neighbours
func getLowPoints(h heightmap) []point {
	lowestPoints := []point{}
	for row := 0; row < len(h); row++ {
		for col := 0; col < len(h[row]); col++ {
			// check all neighbours
			ns := neighbours(point{row, col})

			thisPoint := point{row, col}
			thisValue := h[row][col]
			isLowest := all(ns, func(p point) bool {
				value, err := h.valueAtPoint(p)
				if err != nil {
					return true
				}
				return value > thisValue
			})
			if isLowest {
				lowestPoints = append(lowestPoints, thisPoint)
			}
		}
	}

	return lowestPoints
}

func all(ps []point, pred func(point) bool) bool {
	for _, p := range ps {
		if !pred(p) {
			return false
		}
	}
	return true
}

func neighbours(p point) []point {
	ns := []point{}
	for row := p.row - 1; row <= p.row+1; row++ {
		for col := p.col - 1; col <= p.col+1; col++ {
			if row == p.row && col == p.col {
				continue
			}
			ns = append(ns, point{row, col})
		}
	}
	return ns
}
