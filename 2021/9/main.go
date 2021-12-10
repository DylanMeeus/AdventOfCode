package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type heightmap [][]int

func (h heightmap) print() {
	for row := 0; row < len(h); row++ {
		for _, cell := range h[row] {
			fmt.Printf("%v\t", cell)
		}
		fmt.Println()
	}
}

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
	fmt.Printf("%v\n", solve2())
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

func solve2() int {
	data := getData()

	identifier := -1
	basinSize := map[int]int{}
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			floodFill(data, point{row, col}, identifier, basinSize)
			identifier--
		}
		identifier--
	}

	values := sortMapValues(basinSize)

	mult := 1
	for _, v := range values[:3] {
		mult *= v
	}

	return mult
}

func sortMapValues(m map[int]int) []int {
	out := []int{}
	for _, v := range m {
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] > out[j] })
	return out
}

func floodFill(data heightmap, p point, identifier int, basinSize map[int]int) {
	if value, err := data.valueAtPoint(p); err == nil {
		if value == 9 {
			return
		}
		if value < 0 {
			// this has already been explored
			return
		}
		data[p.row][p.col] = identifier
		basinSize[identifier]++
	} else {
		return
	}
	ns := newNeighbours(p)
	for _, neigh := range ns {
		if val, err := data.valueAtPoint(neigh); err != nil {
			if val == 9 {
				continue
			}
		}
		floodFill(data, neigh, identifier, basinSize)
	}
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

func newNeighbours(p point) []point {
	ns := []point{}
	ns = append(ns, point{p.row - 1, p.col})
	ns = append(ns, point{p.row + 1, p.col})
	ns = append(ns, point{p.row, p.col - 1})
	ns = append(ns, point{p.row, p.col + 1})
	return ns
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
