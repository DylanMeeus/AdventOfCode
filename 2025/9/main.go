package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	NORTHMOST = 0
)

type Point struct {
	row, col int
}

func main() {
	ps := parse(readInput())
	//fmt.Println(solve1(ps))
	northMostVal(ps)
	fmt.Println(solve2(ps))
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

func drawLine(p1, p2 Point) []Point {
	// same row
	out := []Point{}
	if p1.row == p2.row {
		// do we move up or down?
		if p1.col <= p2.col {
			for i := p1.col; i < p2.col; i++ {
				out = append(out, Point{p1.row, i})
			}
		}
		if p1.col > p2.col {
			for i := p1.col; i > p2.col; i-- {
				out = append(out, Point{p1.row, i})
			}
		}
	} else if p1.col == p2.col {
		if p1.row <= p2.row {
			for i := p1.row; i < p2.row; i++ {
				out = append(out, Point{i, p1.col})
			}
		}
		if p1.row > p2.row {
			for i := p1.row; i > p2.row; i-- {
				out = append(out, Point{i, p1.col})
			}
		}
	} else {
		fmt.Printf("%v %v\n", p1, p2)
		panic("This should not happen yo")
	}
	return out
}

func toMap(ps []Point) map[Point]bool {
	m := map[Point]bool{}
	for _, p := range ps {
		m[p] = true
	}
	return m
}

var memo = map[Point]bool{}

// return the minimum row encountered
func northMostVal(ps []Point) int {
	row := ps[0].row
	for _, p := range ps {
		if p.row < row {
			row = p.row
		}
	}
	return row
}

func inBoundary(p Point, m map[Point]bool) bool {
	WINDOW := 10000
	// it's in the figure if in each direction it only intersects _ONE_ boundary

	if m[p] {
		return true
	}

	if val, ok := memo[p]; ok {
		return val
	}

	encountered := []Point{}

	start := p.row
	lookAhead := Point{row: start - 1, col: p.col}
	if m[p] && m[lookAhead] { // this means it lies on the edge
		return true
	}
	intersectsNorth := 0
	for i := 0; i < WINDOW; i++ {
		newPoint := Point{row: start, col: p.col}
		encountered = append(encountered, newPoint)
		lookAhead := Point{row: start - 1, col: p.col}
		if m[newPoint] && !m[lookAhead] {
			intersectsNorth++
		}
		start--
		if start < NORTHMOST {
			break
		}
	}
	if intersectsNorth == 0 || intersectsNorth%2 == 0 {
		for _, enc := range encountered {
			memo[enc] = false
		}
		return false
	}
	for _, enc := range encountered {
		memo[enc] = true
	}
	encountered = nil
	return true
}

type Pair struct {
	a, b Point
}

func sortedPairs(points []Point) []Pair {
	pairs := []Pair{}
	for _, p := range points {
		for _, o := range points {
			if p != o {
				pairs = append(pairs, Pair{p, o})
			}
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return area(pairs[i].a, pairs[i].b) >= area(pairs[j].a, pairs[j].b) })

	return pairs
}

// floodFill the entire rectangle so it's easier to do a 'contain' check
func floodFill(ps []Point) map[Point]bool {
	return nil
}

func solve2(points []Point) int {
	// draw a polygon essentially
	// then 'solve1' with points that are enclosed within this shape..
	// close the loop
	points = append(points, points[0])

	border := []Point{}

	for i := 0; i < len(points)-1; i++ {
		current := points[i]
		next := points[i+1]
		newPoints := drawLine(current, next)
		border = append(border, newPoints...)
	}
	//draw(border)

	borderMap := toMap(border)

	containsAll := func(p1, p2, p3, p4 Point) bool {
		return inBoundary(p1, borderMap) && inBoundary(p2, borderMap) && inBoundary(p3, borderMap) && inBoundary(p4, borderMap)
	}

	// now we need to select 2 points that do not paint outside this figure

	pairs := sortedPairs(points)

	LENGTH := len(pairs)
outer:
	for i, pair := range pairs {
		fmt.Printf("checking pair %v of %v\n", i, LENGTH)
		point := pair.a
		other := pair.b
		if point == other {
			continue
		}
		// the points need to be on a diagonal
		if point.row == other.row || point.col == other.col {
			continue
		}

		// we need to create 2 points to close the loop

		// figure out the correct bounding box

		var x, y Point

		x = Point{row: other.row, col: point.col}
		y = Point{row: point.row, col: other.col}

		lines := []Point{point, x, other, y, point}
		rect := []Point{}
		for i := 0; i < len(lines)-1; i++ {
			pts := drawLine(lines[i], lines[i+1])
			rect = append(rect, pts...)
		}

		if !containsAll(point, other, x, y) {
			continue outer
		}

		a := area(point, other)
		valid := true
		for _, l := range rect {
			if !inBoundary(l, borderMap) {
				valid = false
				continue outer
			}
		}

		if valid {
			// return the first valid area
			return a
		}
	}

	// now for each point,figure out if it is in the border
	//draw(border)

	return -1
}

func parse(lines []string) []Point {
	out := []Point{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		coli, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		rowi, err := strconv.Atoi(parts[1])
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

func draw(ps []Point) {
	m := map[Point]bool{}

	for _, p := range ps {
		m[p] = true
	}

	for row := 0; row < 13; row++ {
		for col := 0; col < 13; col++ {
			if _, ok := m[Point{row, col}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
