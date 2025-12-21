package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point struct {
	row, col int
}

func main() {
	m := toPointMap(readInput())
	//fmt.Println(solve1(m))
	fmt.Println(solve2(m))
}

func solve1(m map[Point]rune) int {

	count := 0
	for key, val := range m {
		if val == '@' {
			c := countNeighbours(key, m)
			if c < 4 {
				count++
			}
		}
	}

	return count

}

func solve2(m map[Point]rune) int {
	// solve2 figures out how many roles of toilet paper we can remove

	removed := 0
	marked := []Point{}
loop:
	for key, val := range m {
		if val == '@' {
			c := countNeighbours(key, m)
			if c < 4 {
				marked = append(marked, key)
			}
		}
	}
	// wipe all of the ones found.
	if len(marked) > 0 {
		removed += len(marked)
		for _, p := range marked {
			m[p] = '.'
		}
		marked = []Point{}
		goto loop
	}

	return removed
}

func countNeighbours(p Point, m map[Point]rune) int {
	directions := []Point{
		// top row
		Point{p.row - 1, p.col - 1},
		Point{p.row - 1, p.col},
		Point{p.row - 1, p.col + 1},
		// middle row
		Point{p.row, p.col - 1},
		Point{p.row, p.col + 1},
		// bottom row
		Point{p.row + 1, p.col - 1},
		Point{p.row + 1, p.col},
		Point{p.row + 1, p.col + 1},
	}
	count := 0
	for _, dir := range directions {
		if m[dir] == '@' {
			count++
		}
		// early exit for perf
		if count == 4 {
			return count
		}
	}
	return count
}

func toPointMap(lines []string) map[Point]rune {
	out := map[Point]rune{}
	for row, line := range lines {
		for col, b := range line {
			loc := Point{row, col}
			out[loc] = b
		}
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
