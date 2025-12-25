package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	row, col int
}

func main() {
	ps := parse(readInput())
	fmt.Println(solve1(ps))
}

func area(p1, p2 Point) int {

	Lf := float64((p1.row - p2.row) + 1)
	Hf := float64(p1.col-p2.col) + 1
	H := int(math.Abs(Lf))
	L := int(math.Abs(Hf))

	return H * L
}

func solve1(points []Point) int {
	maxArea := 0

	for _, point := range points {
		for _, other := range points {
			if point == other {
				continue
			}
			// else figure out the size
			a := area(point, other)
			if a > maxArea {
				maxArea = a
			}
		}
	}

	return maxArea
}

func parse(lines []string) []Point {
	out := []Point{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		rowi, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		coli, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		out = append(out, Point{rowi, coli})
	}
	return out
}

func readInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
