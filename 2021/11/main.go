package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type octopi [][]int

type point struct{ row, col int }

// flash increases the octopis energy level at point p
func (o octopi) levelUp(p point, hasFlashed map[point]bool) {
	if p.row < 0 || p.row >= len(o) {
		return
	}
	if p.col < 0 || p.col >= len(o) {
		return
	}

	if hasFlashed[p] {
		return
	}

	o[p.row][p.col]++
	if o[p.row][p.col] > 9 {
		o.flash(p, hasFlashed)
	}
}

func (o octopi) flash(p point, hasFlashed map[point]bool) {
	// flash this one (reset value to 0, flash all neighbours)
	o[p.row][p.col] = 0
	hasFlashed[p] = true
	for _, n := range neighbours(p) {
		o.levelUp(n, hasFlashed)
	}
}

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

func getData() octopi {
	in, _ := ioutil.ReadFile("./input.txt")

	matrix := [][]int{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		colls := []int{}

		for _, char := range strings.Split(line, "") {
			i, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			colls = append(colls, i)
		}

		matrix = append(matrix, colls)
	}

	return octopi(matrix)
}

func solve() int {
	data := getData()

	MAX_ITER := 100
	flashCount := 0
	for i := 0; i < MAX_ITER; i++ {
		flashSet := map[point]bool{}
		for row := 0; row < len(data); row++ {
			for col := 0; col < len(data); col++ {
				p := point{row, col}
				data.levelUp(p, flashSet)
			}
		}
		flashCount += len(flashSet)
	}

	return flashCount
}

func solve2() int {
	data := getData()

	step := 0
	flashSet := map[point]bool{}
	for len(flashSet) != 100 {
		flashSet = map[point]bool{}
		for row := 0; row < len(data); row++ {
			for col := 0; col < len(data); col++ {
				p := point{row, col}
				data.levelUp(p, flashSet)
			}
		}
		step++
		fmt.Println(step)
	}

	return step
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
