package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Beam struct {
	location Point
	alive    bool
}

type Point struct {
	row, col int
}

func main() {
	lines := readInput()
	g := len(lines) // length of grid to know when to stop scanning
	s, m := parse(lines)
	//fmt.Println(solve1(s, m, g))
	fmt.Println(solve2(s, m, g))
}

func solve2(start Point, m map[Point]rune, g int) int {
	// the splitters we have hit

	// at this point you do need memoization

	// memoize the 'worlds' after each point
	memo := map[Point]int{}

	beam := Beam{location: start, alive: true}

	var traverse func(Beam, int) int
	traverse = func(b Beam, current int) int {
		// else we update the location
		for m[b.location] != '^' {
			b.location.row++
			if b.location.row > g {
				return current
			}
		}
		if val, ok := memo[b.location]; ok {
			return val
		}
		if m[b.location] == '^' {
			current++
			current += traverse(Beam{location: Point{b.location.row, b.location.col - 1}}, 0)
			current += traverse(Beam{location: Point{b.location.row, b.location.col + 1}}, 0)
			memo[b.location] = current
		}
		return current
	}

	res := traverse(beam, 1)

	return res

}

func solve1(start Point, m map[Point]rune, g int) int {
	// the splitters we have hit
	splitters := map[Point]bool{}

	beams := []*Beam{&Beam{location: start, alive: true}}
	livingBeams := 1

	tick := func() {
		// move it down by one
		for _, beam := range beams {
			if !beam.alive {
				continue
			}
			beam.location.row++
			if m[beam.location] == '^' {
				splitters[beam.location] = true
				beam.alive = false
				livingBeams--
				left := &Beam{location: Point{beam.location.row, beam.location.col - 1}, alive: true}
				if m[left.location] == '.' {
					beams = append(beams, left)
					m[left.location] = '|'
					livingBeams++
				}
				right := &Beam{location: Point{beam.location.row, beam.location.col + 1}, alive: true}
				if m[right.location] == '.' {
					beams = append(beams, right)
					livingBeams++
					m[right.location] = '|'
				}

				// only if we added something
				// only add a beam if none yet exists at this location?
			} else {
				m[beam.location] = '|'
			}
			if beam.location.row > g {
				beam.alive = false
				livingBeams--
			}
		}
	}
	// while we have beams that are alive..

	for livingBeams != 0 {
		tick()
	}

	return len(splitters)
}

func parse(lines []string) (Point, map[Point]rune) {
	m := map[Point]rune{}
	s := Point{}
	for row, line := range lines {
		for col, char := range line {
			p := Point{row, col}
			m[p] = char
			if char == 'S' {
				s = p
			}

		}
	}
	return s, m
}

func readInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func print(m map[Point]rune) {
	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			fmt.Print(string(m[Point{i, j}]))
		}
		fmt.Println()
	}
}
