package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve())
	solve2()
}

type direction int

const (
	X direction = iota
	Y
)

type point struct{ x, y int }

type fold struct {
	d        direction
	position int
}

type data struct {
	points []point
	folds  []fold
}

func getData() data {
	in, _ := ioutil.ReadFile("./input.txt")

	folds := []fold{}
	points := []point{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold") {
			parts := strings.Split(line, " ")
			pos := parts[len(parts)-1]
			foldParts := strings.Split(pos, "=")

			dir := X
			if foldParts[0] == "y" {
				dir = Y
			}

			position, err := strconv.Atoi(foldParts[1])
			if err != nil {
				panic(err)
			}
			folds = append(folds, fold{dir, position})
		} else {
			coords := strings.Split(line, ",")
			x, err := strconv.Atoi(coords[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				panic(err)
			}
			points = append(points, point{x, y})
		}
	}

	return data{points, folds}
}

func solve() int {
	input := getData()

	for _, fold := range input.folds[0:1] {
		for i, point := range input.points {
			if fold.d == X {
				input.points[i] = foldLeft(fold.position, point)
			} else {
				input.points[i] = foldUp(fold.position, point)
			}
		}
	}

	// now count unique points ("visible")

	m := map[point]bool{}
	for _, point := range input.points {
		if point.x < 0 || point.y < 0 {
			continue
		}
		m[point] = true
	}

	return len(m)
}

func solve2() {
	input := getData()

	for _, fold := range input.folds {
		for i, point := range input.points {
			if fold.d == X {
				input.points[i] = foldLeft(fold.position, point)
			} else {
				input.points[i] = foldUp(fold.position, point)
			}
		}
	}

	// now count unique points ("visible")

	m := map[point]bool{}
	for _, point := range input.points {
		if point.x < 0 || point.y < 0 {
			continue
		}
		m[point] = true
	}

	plot(m)

	fmt.Printf("%v\n", input.folds)
}

func plot(m map[point]bool) {

	rows := []int{}
	columns := []int{}

	for p, _ := range m {
		rows = append(rows, p.y)
		columns = append(columns, p.x)
	}

	maxRows := max(rows)
	maxColumns := max(columns)

	for row := 0; row <= maxRows; row++ {
		for col := 0; col <= maxColumns; col++ {
			if m[point{col, row}] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n\n")
	}
}

func max(is []int) int {
	res := 0
	for _, i := range is {
		if i > res {
			res = i
		}
	}
	return res
}

func foldUp(y int, p point) point {
	if p.y <= y {
		return p
	}
	delta := p.y - y
	p.y = y - delta
	return p
}

func foldLeft(x int, p point) point {
	if p.x <= x {
		return p
	}
	delta := p.x - x
	p.x = x - delta
	return p
}
